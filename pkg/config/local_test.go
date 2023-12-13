package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadLocalConfig(t *testing.T) {
	configDir := "."
	configName := "test.config.yaml"
	expectedConfig := &Config{
		JiraUrl: "test.url",
		Token: "token",
		Project: "proj",
	}

	result, err := LoadLocalConfig(configDir, configName)

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, result)
}
