package generator

import (
	"encoding/json"
	"os"
	"path/filepath"

	file "github.com/dterbah/go-test-gen/utils"
	"github.com/sirupsen/logrus"
)

const GO_TEST_GEN_CONFIG_NAME = ".go-gen-test.json"

type TestGeneratorConfig struct {
	GeneratePrivateFunctions bool     `json:"generatePrivateFunctions"`
	GenerateEmptyTests       bool     `json:"generateEmptyTests"`
	ExcludeFiles             []string `json:"excludeFiles"`
}

var defaultConfig TestGeneratorConfig = TestGeneratorConfig{
	GeneratePrivateFunctions: false,
	GenerateEmptyTests:       false,
	ExcludeFiles:             []string{},
}

func LoadTestGeneratorConfig(path string) TestGeneratorConfig {
	configPath := filepath.Join(path, GO_TEST_GEN_CONFIG_NAME)
	if file.Exists(configPath) {
		logrus.Infof("🟢 Configuration file found !")
		file, err := os.Open(configPath)
		if err != nil {
			logrus.Error("Error when opening the config file")
			return defaultConfig
		}

		defer file.Close()

		decoder := json.NewDecoder(file)
		config := TestGeneratorConfig{}
		err = decoder.Decode(&config)
		if err != nil {
			logrus.Errorf("Error when decoding the config file <%s>", err)
			return defaultConfig
		}

		logrus.Info(config.GeneratePrivateFunctions)
		return config
	}

	logrus.Infof("🟠 No config file found. Take default config instead")
	return defaultConfig
}
