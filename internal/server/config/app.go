package config

import (
	"context"
	"os"

	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/viper"
	"github.com/status-mok/server/internal/pkg/errors"
	"github.com/status-mok/server/internal/pkg/log"
	"go.uber.org/zap/zapcore"
)

type AppConfig struct {
	LogLevelStr string `mapstructure:"log_level" env:"MOK_LOG_LEVEL,default=debug"`
	AdminAPI    struct {
		GRPC struct {
			Host string `mapstructure:"host" env:"MOK_ADMIN_API_GRPC_HOST"`
			Port string `mapstructure:"port" env:"MOK_ADMIN_API_GRPC_PORT,default=2001"`
		} `mapstructure:"grpc"`
		HTTP struct {
			Host string `mapstructure:"host" env:"MOK_ADMIN_API_HTTP_HOST"`
			Port string `mapstructure:"port" env:"MOK_ADMIN_API_HTTP_PORT,default=2002"`
		} `mapstructure:"http"`
	} `mapstructure:"admin_api"`
}

func NewAppConfig(ctx context.Context, configPath string) (*AppConfig, error) {
	var conf AppConfig
	if len(configPath) == 0 {
		if err := envconfig.Process(ctx, &conf); err != nil {
			return nil, errors.Wrap(err, "failed to process config environment variables")
		}
		return &conf, nil
	}

	f, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open config file '%s'", configPath)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(f); err != nil {
		return nil, errors.Wrap(err, "failed to read config yaml file")
	}
	if err := v.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "failed to decode config yaml file")
	}

	return &conf, nil
}

func (conf AppConfig) LogLevel() zapcore.LevelEnabler {
	level, err := zapcore.ParseLevel(conf.LogLevelStr)
	if err != nil {
		return log.DefaultLevel
	}

	return level
}

func (conf AppConfig) AdminHTTPAddress() string {
	return conf.AdminAPI.HTTP.Host + ":" + conf.AdminAPI.HTTP.Port
}

func (conf AppConfig) AdminGRPCAddress() string {
	return conf.AdminAPI.GRPC.Host + ":" + conf.AdminAPI.GRPC.Port
}
