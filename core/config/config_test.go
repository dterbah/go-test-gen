package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeConfigurationHelper(config []byte, missingPermission bool) (string, error) {
	tmpFile, err := os.Create(".go-test-gen.json")
	if err != nil {
		return "", err
	}

	if _, err := tmpFile.Write(config); err != nil {
		return "", err
	}
	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	if missingPermission {
		os.Chmod(tmpFile.Name(), 0000)
	}

	return tmpFile.Name(), nil
}

func TestLoadTestGeneratorConfigWithGeneratePrivateFunctions(t *testing.T) {
	assert := assert.New(t)
	configContent := []byte(`
		{
  			"generatePrivateFunctions": true
		}`)

	filename, err := writeConfigurationHelper(configContent, false)
	if err != nil {
		assert.FailNow(err.Error())
	}

	defer os.Remove(filename)

	config := LoadTestGeneratorConfig(".")

	assert.True(config.GeneratePrivateFunctions)
	assert.False(config.GenerateEmptyTests)
	assert.Equal([]string{}, config.ExcludeFiles)
	assert.True(config.Verbose)
}

func TestLoadTestGeneratorConfigWithExcludeFiles(t *testing.T) {
	assert := assert.New(t)
	configContent := []byte(`
		{
  			"excludeFiles": ["main.go", "toto.go"]
		}`)
	filename, err := writeConfigurationHelper(configContent, false)
	if err != nil {
		assert.FailNow(err.Error())
	}

	defer os.Remove(filename)

	config := LoadTestGeneratorConfig(".")
	expectedExcludedFiles := []string{"main.go", "toto.go"}
	assert.Equal(expectedExcludedFiles, config.ExcludeFiles)
	assert.False(config.GeneratePrivateFunctions)
	assert.True(config.Verbose)
}

func TestLoadTestGeneratorConfigWithVerbose(t *testing.T) {
	assert := assert.New(t)
	configContent := []byte(`
		{
  			"verbose": false
		}`)
	filename, err := writeConfigurationHelper(configContent, false)
	if err != nil {
		assert.FailNow(err.Error())
	}

	defer os.Remove(filename)

	config := LoadTestGeneratorConfig(".")

	assert.False(config.Verbose)
}

func TestLoadTestGeneratorConfigWithFile(t *testing.T) {
	assert := assert.New(t)

	config := LoadTestGeneratorConfig(".")
	assert.Equal(config, defaultConfig)
}

func TestLoadTestGeneratorConfigWithMissingPermission(t *testing.T) {
	assert := assert.New(t)
	configContent := []byte(`
		{
  			"verbose": false
		}`)
	filename, err := writeConfigurationHelper(configContent, true)
	if err != nil {
		assert.FailNow(err.Error())
	}

	defer os.Remove(filename)

	config := LoadTestGeneratorConfig(".")
	assert.Equal(config, defaultConfig)
}

func TestLoadTestGeneratorConfigWithBadConfig(t *testing.T) {
	assert := assert.New(t)
	configContent := []byte(`
		{
  			"verbo": false
		`)
	filename, err := writeConfigurationHelper(configContent, false)
	if err != nil {
		assert.FailNow(err.Error())
	}

	defer os.Remove(filename)

	config := LoadTestGeneratorConfig(".")
	assert.Equal(config, defaultConfig)
}
