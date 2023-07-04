package logger

import (
	"os"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	log_level "github.com/y-du/go-log-level"
)

var Logger *log_level.Logger

func InitLogger(config srv_base.LoggerConfig) (out *os.File, err error) {
	Logger, out, err = srv_base.NewLogger(config)
	return
}
