package config

import (
	"io/fs"
	"os"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/y-du/go-log-level/level"
)

type SocketConfig struct {
	Path     string      `json:"path" env_var:"SOCKET_PATH"`
	GroupID  int         `json:"group_id" env_var:"SOCKET_GROUP_ID"`
	FileMode fs.FileMode `json:"file_mode" env_var:"SOCKET_FILE_MODE"`
}
type Config struct {
	Logger           srv_base.LoggerConfig `json:"logger" env_var:"LOGGER_CONFIG"`
	TMPFSPath        string                `json:"tmpfs_path" env_var:"TMPFS_PATH"`
	EnableEncryption bool                  `json:"enable_encryption" env_var:"ENABLE_ENCRYPTION"`
	DBFilePath       string                `json:"db_file_path" env_var:"DB_FILE_PATH"`
	MasterKeyPath    string                `json:"master_key_file_path" env_var:"MASTER_KEY_FILE_PATH"`
	Socket           SocketConfig          `json:"socket" env_var:"SOCKET_CONFIG"`
}

func NewConfig(path *string) (*Config, error) {
	cfg := Config{
		Logger: srv_base.LoggerConfig{
			Level:        level.Debug,
			Utc:          true,
			Path:         ".",
			FileName:     "github.com/SENERGY-Platform/mgw-secret-manager",
			Microseconds: true,
		},
		Socket: SocketConfig{
			Path:     "./sm.sock",
			GroupID:  os.Getgid(),
			FileMode: 0660,
		},
		TMPFSPath:        "/tmp",
		EnableEncryption: true,
		DBFilePath:       "./db.sqlite",
		MasterKeyPath:    "./key",
	}
	if *path == "" {
		path = nil
	}
	err := srv_base.LoadConfig(path, &cfg, nil, nil, nil)
	return &cfg, err
}
