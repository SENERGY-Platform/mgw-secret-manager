package main

import (
	"context"
	"errors"
	"fmt"
	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	sb_logger "github.com/SENERGY-Platform/go-service-base/logger"
	srv_info_hdl "github.com/SENERGY-Platform/go-service-base/srv-info-hdl"
	sb_util "github.com/SENERGY-Platform/go-service-base/util"
	"github.com/SENERGY-Platform/go-service-base/watchdog"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/key_handler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secret_handler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"
)

var version string

func main() {
	srvInfoHdl := srv_info_hdl.New("secret-manager", version)

	ec := 0
	defer func() {
		os.Exit(ec)
	}()

	util.ParseFlags()

	config, err := util.NewConfig(util.Flags.ConfPath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		ec = 1
		return
	}

	logFile, err := util.InitLogger(config.Logger)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		var logFileError *sb_logger.LogFileError
		if errors.As(err, &logFileError) {
			ec = 1
			return
		}
	}
	if logFile != nil {
		defer logFile.Close()
	}

	util.Logger.Printf("%s %s", srvInfoHdl.GetName(), srvInfoHdl.GetVersion())

	util.Logger.Debugf("config: %s", sb_util.ToJsonStr(config))

	watchdog.Logger = util.Logger
	wtchdg := watchdog.New(syscall.SIGINT, syscall.SIGTERM)

	dbHandler, err := db.NewDBHandler(config)
	if err != nil {
		util.Logger.Error(err)
		ec = 1
		return
	}

	secretHdl := secret_handler.NewSecretHandler(config.EnableEncryption, dbHandler, config.TMPFSPath)

	keyHdl := key_handler.NewKeyHandler(config.MasterKeyPath, nil)

	mApi := api.New(*config, dbHandler, secretHdl, keyHdl, srvInfoHdl)

	gin.SetMode(gin.ReleaseMode)
	httpHandler := gin.New()
	staticHeader := map[string]string{
		http_handler.HeaderApiVer:  srvInfoHdl.GetVersion(),
		http_handler.HeaderSrvName: srvInfoHdl.GetName(),
	}
	httpHandler.Use(gin_mw.StaticHeaderHandler(staticHeader), requestid.New(requestid.WithCustomHeaderStrKey(http_handler.HeaderRequestID)), gin_mw.LoggerHandler(util.Logger, http_handler.GetPathFilter(), func(gc *gin.Context) string {
		return requestid.Get(gc)
	}), gin_mw.ErrorHandler(http_handler.GetStatusCode, ", "), gin.Recovery())
	httpHandler.UseRawPath = true

	http_handler.SetRoutes(httpHandler, mApi)

	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(config.ServerPort, 10))
	if err != nil {
		util.Logger.Error(err)
		ec = 1
		return
	}
	server := &http.Server{Handler: httpHandler}
	srvCtx, srvCF := context.WithCancel(context.Background())
	wtchdg.RegisterStopFunc(func() error {
		if srvCtx.Err() == nil {
			ctxWt, cf := context.WithTimeout(context.Background(), time.Second*5)
			defer cf()
			if err := server.Shutdown(ctxWt); err != nil {
				return err
			}
			util.Logger.Info("http server shutdown complete")
		}
		return nil
	})
	wtchdg.RegisterHealthFunc(func() bool {
		if srvCtx.Err() == nil {
			return true
		}
		util.Logger.Error("http server closed unexpectedly")
		return false
	})

	wtchdg.Start()

	go func() {
		defer srvCF()
		util.Logger.Info("starting http server ...")
		if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
			util.Logger.Error(err)
			ec = 1
			return
		}
	}()

	ec = wtchdg.Join()
}
