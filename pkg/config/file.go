package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/mishamyrt/checode/v1/pkg/types"
	"gopkg.in/yaml.v2"
)

const (
	// ErrKeyword key.
	ErrKeyword = "err"
	// WarnKeyword key.
	WarnKeyword = "warn"
)

func parseFlag(value string) uint8 {
	switch value {
	case ErrKeyword:
		return ErrFlag
	case WarnKeyword:
		return WarnFlag
	default:
		return WarnFlag
	}
}

// ParseFlags transforms text map to bit values.
func ParseFlags(config *types.Config, configMap *types.KewordsMap) {
	for key := range *configMap {
		(*config)[key] = parseFlag((*configMap)[key])
	}
}

// ReadConfigFile reads the configuration from the yaml file.
func ReadConfigFile(path string, config *types.Config) (err error) {
	if fileExists(path) {
		var configFile types.ConfigFile
		err = readYamlConfig(path, &configFile)
		if err != nil {
			return
		}
		ParseFlags(config, &configFile.Keywords)
	} else {
		err = errors.New("file doesn't exist")
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
