package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	srv_base_types "github.com/SENERGY-Platform/go-service-base/srv-base/types"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
)

var version string

func main() {
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

	logger.Logger.Debugf("config: %s", srv_base.ToJsonStr(config))

	httpHandler, _, _ := api.InitServer(config)
	listener, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(config.ServerPort), 10))
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	srv_base.StartServer(&http.Server{Handler: httpHandler}, listener, srv_base_types.DefaultShutdownSignals, logger.Logger)

}
