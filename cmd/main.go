package main

import (
	"errors"
	"fmt"
	"os"
	"secret-manager/internal/api"
	"secret-manager/internal/config"
	"secret-manager/internal/db"

	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/gin-gonic/gin"
)

var version string

func main() {
	srv_base.PrintInfo("mgw-secret-manager", version)

	flags := config.NewFlags()
	config, err := config.NewConfig(flags.ConfPath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	logFile, err := srv_base.InitLogger(config.Logger)
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

	srv_base.Logger.Debugf("config: %s", srv_base.ToJsonStr(config))

	dbHandler, err := db.NewDBHandler(*config)
	if err != nil {
		srv_base.Logger.Error(err)
	}

	gin.SetMode(gin.ReleaseMode)
	apiEngine := gin.New()
	apiEngine.Use(gin_mw.LoggerHandler(srv_base.Logger), gin_mw.ErrorHandler, gin.Recovery())
	apiEngine.UseRawPath = true
	Api := api.New(*config, dbHandler)
	Api.SetRoutes(apiEngine)

	/*listener, err := srv_base.NewUnixListener(config.Socket.Path, os.Getuid(), config.Socket.GroupID, config.Socket.FileMode)
	if err != nil {
		srv_base.Logger.Error(err)
		return
	}*/

	apiEngine.Run()
	//srv_base.StartServer(&http.Server{Handler: apiEngine}, listener, srv_base_types.DefaultShutdownSignals)
}
