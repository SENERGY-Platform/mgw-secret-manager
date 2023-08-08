package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/server"
)

var version string

func main() {
	ec := 0
	defer func() {
		os.Exit(ec)
	}()

	srv_base.PrintInfo("mgw-github.com/SENERGY-Platform/mgw-secret-manager", version)

	config.ParseFlags()

	config, err := config.NewConfig(config.Flags.ConfPath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	logFile, err := logger.InitLogger(config.Logger)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		var logFileError *srv_base.LogFileError
		if errors.As(err, &logFileError) {
			os.Exit(1)
		}
	}
	if logFile != nil {
		defer logFile.Close()
	}

	watchdog := srv_base.NewWatchdog(logger.Logger, syscall.SIGINT, syscall.SIGTERM)

	logger.Logger.Debugf("config: %s", srv_base.ToJsonStr(config))

	httpHandler, _, _ := server.InitServer(config)
	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(config.ServerPort), 10))
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	server := &http.Server{Handler: httpHandler}

	srvCtx, srvCF := context.WithCancel(context.Background())
	watchdog.RegisterStopFunc(func() error {
		if srvCtx.Err() == nil {
			ctxWt, cf := context.WithTimeout(context.Background(), time.Second*5)
			defer cf()
			if err := server.Shutdown(ctxWt); err != nil {
				return err
			}
			logger.Logger.Info("http server shutdown complete")
		}
		return nil
	})

	watchdog.Start()

	go func() {
		defer srvCF()
		logger.Logger.Info("starting http server ...")
		if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Error(err)
			ec = 1
			return
		}
	}()

	ec = watchdog.Join()
}
