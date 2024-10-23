package util

import (
	"github.com/SENERGY-Platform/go-service-base/config-hdl"
	sb_logger "github.com/SENERGY-Platform/go-service-base/logger"
	"github.com/y-du/go-env-loader"
	"github.com/y-du/go-log-level/level"
	"reflect"
)

type LoggerConfig struct {
	Level        level.Level `json:"level" env_var:"LOGGER_LEVEL"`
	Utc          bool        `json:"utc" env_var:"LOGGER_UTC"`
	Path         string      `json:"path" env_var:"LOGGER_PATH"`
	FileName     string      `json:"file_name" env_var:"LOGGER_FILE_NAME"`
	Terminal     bool        `json:"terminal" env_var:"LOGGER_TERMINAL"`
	Microseconds bool        `json:"microseconds" env_var:"LOGGER_MICROSECONDS"`
	Prefix       string      `json:"prefix" env_var:"LOGGER_PREFIX"`
}

type Config struct {
	Logger                      LoggerConfig `json:"logger" env_var:"LOGGER_CONFIG"`
	TMPFSPath                   string       `json:"tmpfs_path" env_var:"TMPFS_PATH"`
	EnableEncryption            bool         `json:"enable_encryption" env_var:"ENABLE_ENCRYPTION"`
	DBConnectionURL             string       `json:"db_connection_url" env_var:"DB_CONNECTION_URL"`
	MasterKeyPath               string       `json:"master_key_file_path" env_var:"MASTER_KEY_FILE_PATH"`
	ServerPort                  int64        `json:"server_port" env_var:"SERVER_PORT"`
	ExposeConfidentialEndpoints bool         `json:"expose_confidential_endpoints" env_var:"EXPOSE_CONFIDENTIAL_ENDPOINTS"`
}

func NewConfig(path string) (*Config, error) {
	cfg := Config{
		Logger: LoggerConfig{
			Level:        level.Warning,
			Utc:          true,
			Microseconds: true,
			Terminal:     true,
		},
		TMPFSPath:                   "/tmp",
		EnableEncryption:            false,
		DBConnectionURL:             "./db.sqlite",
		MasterKeyPath:               "./key",
		ServerPort:                  80,
		ExposeConfidentialEndpoints: false,
	}

	err := config_hdl.Load(&cfg, nil, map[reflect.Type]envldr.Parser{reflect.TypeOf(level.Off): sb_logger.LevelParser}, nil, path)
	return &cfg, err
}
