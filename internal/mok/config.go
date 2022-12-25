package mok

import (
	"io"

	"github.com/spf13/viper"
	"github.com/status-mok/server/internal/pkg/errors"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Servers []ServerConfig `mapstructure:"servers"`
	}

	ServerConfig struct {
		Name      string `mapstructure:"name" yaml:"name"`
		IP        string `mapstructure:"ip" yaml:"ip"`
		Port      uint16 `mapstructure:"port" yaml:"port"`
		Type      string `mapstructure:"type" yaml:"type"`
		IsStopped bool   `mapstructure:"is_stopped" yaml:"is_stopped"`

		Routes []RouteConfig `mapstructure:"routes" yaml:"routes"`
	}

	RouteConfig struct {
		URL        string `mapstructure:"url" yaml:"url"`
		Type       string `mapstructure:"type" yaml:"type"`
		IsDisabled bool   `mapstructure:"is_disabled" yaml:"is_disabled"`

		Expectations []ExpectationConfig `mapstructure:"expectations" yaml:"expectations"`
	}

	ExpectationConfig struct {
		ID         string `mapstructure:"id" yaml:"id"`
		IsDisabled bool   `mapstructure:"is_disabled" yaml:"is_disabled"`
	}
)

func ReadConfig(r io.Reader) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(r); err != nil {
		return nil, errors.Wrap(err, "failed to read config yaml file")
	}

	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "failed to decode config yaml file")
	}

	return &conf, nil
}

func (conf *Config) DumpConfig() ([]byte, error) {
	return yaml.Marshal(conf)
}
