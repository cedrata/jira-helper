package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadLocalConfig(t *testing.T) {

	t.Run("File not found", func(t *testing.T) {
		t.Parallel()

		resultViper := viper.New()
		invalidConfigPath := "/foo"
		err := LoadLocalConfig(invalidConfigPath, DefaultConfigName, resultViper)

		assert.Error(t, err,
			fmt.Sprintf("The configuration reading should fail because not such file as %s exists", invalidConfigPath),
		)

		_, ok := err.(viper.ConfigFileNotFoundError)
		assert.True(t, ok,
			"The error should be of type viper.CofigFileNotFoundError",
		)
	})

	t.Run("Wrong file format", func(t *testing.T) {
		t.Parallel()

		resultViper := viper.New()
		configDir := t.TempDir()
		configContent := "[default]\nurl=test.url\ntoken=token\nfoo"
		configFullPath := filepath.Join(configDir, DefaultConfigName)

		_ = os.WriteFile(configFullPath, []byte(configContent), 0777)

		err := LoadLocalConfig(configDir, DefaultConfigName, resultViper)

		assert.Error(t, err,
			"Reading an invalid formatted configuration should result in error",
		)

		_, ok := err.(viper.ConfigFileNotFoundError)
		assert.False(t, ok,
			"The error should not be of type viper.CofigFileNotFoundError",
		)
	})

	t.Run("Load Ok", func(t *testing.T) {
		t.Parallel()

		resultViper := viper.New()
		configDir := t.TempDir()
		configContent := "[default]\nurl=value1\ntoken=value2"
		configFullPath := filepath.Join(configDir, DefaultConfigName)

		_ = os.WriteFile(configFullPath, []byte(configContent), 0777)

		err := LoadLocalConfig(configDir, DefaultConfigName, resultViper)
		assert.NoError(t, err)

		expectedContent := map[string]string{
			"url":   "value1",
			"token": "value2",
		}

		actualContent := resultViper.GetStringMapString("default")
		assert.Equal(t, expectedContent, actualContent)
	})
}
