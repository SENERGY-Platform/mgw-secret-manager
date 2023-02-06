package test

import (
	"os"
	"secret-manager/internal/config"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/y-du/go-log-level/level"
)

var MasterKey []byte = make([]byte, 32)

var TestConfig config.Config = config.Config{
	Logger: srv_base.LoggerConfig{
		Level:        level.Debug,
		Utc:          true,
		Path:         ".",
		FileName:     "secret-manager",
		Microseconds: true,
	},
	Socket: config.SocketConfig{
		Path:     "./sm.sock",
		GroupID:  os.Getgid(),
		FileMode: 0660,
	},
	DBFilePath:       "test_db.sqlite",
	EnableEncryption: false,
}
