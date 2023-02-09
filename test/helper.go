package test

import (
	"os"
	"secret-manager/internal/config"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/y-du/go-log-level/level"
)

var MasterKey []byte = make([]byte, 32)

var EncryptionKey = []byte("eShVmYq3t6w9z$C&E)H@McQfTjWnZr4u")

var TestConfig config.Config = config.Config{
	Logger: srv_base.LoggerConfig{
		Level:        level.Debug,
		Utc:          true,
		Path:         "/tmp",
		FileName:     "secret-manager",
		Microseconds: true,
	},
	Socket: config.SocketConfig{
		Path:     "./sm.sock",
		GroupID:  os.Getgid(),
		FileMode: 0660,
	},
	DBFilePath:       "/tmp/test_db.sqlite",
	EnableEncryption: true,
	MasterKeyPath:    "/tmp/key",
	TMPFSPath:        "/tmp",
}
