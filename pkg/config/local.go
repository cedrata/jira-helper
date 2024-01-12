package config

import (
	"fmt"
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

func LoadLocalConfig(configFile string, v *viper.Viper) error {
	var err error
    var configPath string
    var configName string

	if strings.TrimSpace(configFile) == "" {
		configPath, err = os.UserHomeDir()
		configName = defaultConfigName
        fmt.Println("from default")
	} else {
		configPath = filepath.Dir(configFile)
		configName = filepath.Base(configFile)
        fmt.Println("from provided")
    }

	if err != nil {
		return err
	}

    v.AddConfigPath(configPath)
    v.SetConfigName(configName)
	v.SetConfigType(configType)
	err = v.ReadInConfig()
    fmt.Println(v.GetString("host"))

	return err
}
