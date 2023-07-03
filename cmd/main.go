package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
)

var version string

func main() {
	srv_base.PrintInfo("mgw-github.com/SENERGY-Platform/mgw-secret-manager", version)

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

	apiEngine, _, _ := api.InitServer(config)
	apiEngine.Run()
	//srv_base.StartServer(&http.Server{Handler: apiEngine}, listener, srv_base_types.DefaultShutdownSignals)
}
