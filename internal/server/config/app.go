package config

import (
	"context"
	"os"

	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	LogLevel         string `mapstructure:"log_level" env:"MOK_LOG_LEVEL,default=error"`
	AdminAPIgrpcHost string `mapstructure:"admin_api_grpc_host" env:"MOK_ADMIN_API_GRPC_HOST"`
	AdminAPIgrpcPort string `mapstructure:"admin_api_grpc_port" env:"MOK_ADMIN_API_GRPC_PORT,default=2001"`
	AdminAPIhttpHost string `mapstructure:"admin_api_http_host" env:"MOK_ADMIN_API_HTTP_HOST"`
	AdminAPIhttpPort string `mapstructure:"admin_api_http_port" env:"MOK_ADMIN_API_HTTP_PORT,default=2002"`
}

func NewAppConfig(ctx context.Context, configPath string) (*AppConfig, error) {
	var conf AppConfig
	if len(configPath) == 0 {
		if err := envconfig.Process(ctx, &conf); err != nil {
			return nil, err
		}
		return &conf, nil
	}

	f, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(f, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (conf AppConfig) AdminHTTPAddress() string {
	return conf.AdminAPIhttpHost + ":" + conf.AdminAPIhttpPort
}

func (conf AppConfig) AdminGRPCAddress() string {
	return conf.AdminAPIgrpcHost + ":" + conf.AdminAPIgrpcPort
}
