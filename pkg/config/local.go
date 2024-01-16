package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	configType = "env"
    defaultConfigName = ".jhelp.config"
)

type Config struct {
	JiraUrl string `mapstructure:"url"`
	Token   string `mapstructure:"token"`
	Project string `mapstructure:"project"`
}

// Loads a configuration file.
// If no configuration file is provided the default one is loaded.
// Returns an error if the file is not available.
func LoadLocalConfig(configFile string, v *viper.Viper) error {
	var err error
    var configPath string
    var configName string

	if strings.TrimSpace(configFile) == "" {
		configPath, err = os.UserHomeDir()
		configName = defaultConfigName
	} else {
		configPath = filepath.Dir(configFile)
		configName = filepath.Base(configFile)
    }

	if err != nil {
		return err
	}

    v.AddConfigPath(configPath)
    v.SetConfigName(configName)
	v.SetConfigType(configType)
	err = v.ReadInConfig()

	return err
}
