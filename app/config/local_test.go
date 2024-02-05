package config

import (
	// "os"
	// "path/filepath"
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

		resultConfig := viper.New()
		invalidConfigPath := "/foo/file.pippo"
		err := LoadLocalConfig(invalidConfigPath, resultConfig)

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

		resultConfig := viper.New()
		configDir := t.TempDir()
		configName := "config.env"
		configContent := "url=test.url\ntoken=token\nproject=proj\nfoo"
		configFullPath := filepath.Join(configDir, configName)

		os.WriteFile(configFullPath, []byte(configContent), 0777)

		err := LoadLocalConfig(configFullPath, resultConfig)

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

		resultConfig := viper.New()
		configDir := t.TempDir()
		configName := "config.env"
		configContent := "url=value1\ntoken=value2\nproject=value3"
		configFullPath := filepath.Join(configDir, configName)

		os.WriteFile(configFullPath, []byte(configContent), 0777)

		err := LoadLocalConfig(configFullPath, resultConfig)

		assert.NoError(t, err)

		expectedKeys := []string{"url", "token", "project"}
		for i, _ := range expectedKeys {
            key := expectedKeys[i]
            value := resultConfig.GetString(key)
            t.Logf("key %s having value: %s", key, value)
			assert.Equal(t, fmt.Sprintf("value%d", i+1), value)
		}
	})
}
