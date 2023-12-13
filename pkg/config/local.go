package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = ".jira.config"
	configDir  = "~"
)

type Config struct {
	JiraUrl string `mapstructure:"url"`
	Token   string `mapstructure:"token"`
	Project string `mapstructure:"project"`
}

func LoadLocalConfig(configDir string, configName string) (*Config, error) {
	var config Config

	v := viper.New()

	v.SetConfigName(configName)
	v.SetConfigType(configType)
	v.AddConfigPath(configDir)

	err := v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, errors.WithStack(err)
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}
