package config

import (
	"github.com/cedrata/jira-helper/pkg/helpers"
	"github.com/spf13/viper"
)

const (
	ConfigType        = "ini"
	DefaultConfigName = ".jhelp.config"
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

// Validate the provided viper instance and if valid return a configuratio struct
func ValidateProfile(profile string, v *viper.Viper) (*Config, error) {
	var err error
	var res Config

	err = v.UnmarshalKey(profile, &res)
	if err != nil {
		return &res, err
	}
    
    err = helpers.ValidateStruct(res)

	return &res, err
}
