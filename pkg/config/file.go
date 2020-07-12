package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/mishamyrt/checode/v1/pkg/types"
	"gopkg.in/yaml.v2"
)

// ReadConfigFile reads the configuration from the yaml file
func ReadConfigFile(path string, config *types.Config) (err error) {
	if fileExists(path) {
		var configFile types.ConfigFile
		err = readYamlConfig(path, &configFile)
		*config = configFile.Keywords
	} else {
		err = errors.New("File doesn't exist")
	}
	return
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readYamlConfig(filePath string, storage *types.ConfigFile) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, storage)
}
