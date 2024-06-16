package config

import (
	"github.com/spf13/viper"
)

const (
	ConfigType        = "toml"
	DefaultConfigName = ".jira-helper.toml"
)

type Config struct {
	Host  string `mapstructure:"host" validation:"required"`
	Token string `mapstructure:"token" validation:"required"`
}

// This is a global configuration variable that should not be modified directly
// other than in cmd/root when reading the configuration.
var ConfigData *Config

// Load the configuration file inside the viper provided instance.
func LoadLocalConfig(configPath string, configName string, v *viper.Viper) error {
	var err error

	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(ConfigType)
	err = v.ReadInConfig()

	return err
}
