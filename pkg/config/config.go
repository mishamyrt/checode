package config

import (
	"io/ioutil"
	"os"

	"github.com/mishamyrt/checode/v1/pkg/types"
	"gopkg.in/yaml.v2"
)

// ConfigPaths is allowed configuration file names
var ConfigPaths = []string{".checode.yml", ".checode.yaml"}

var defaultConfig = types.Keywords{
	"FIXME":    types.ErrKeyword,
	"STOPSHIP": types.ErrKeyword,
	"TODO":     types.WarnKeyword,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func parse(data []byte, storage interface{}) error {
	return yaml.Unmarshal(data, storage)
}

func readYaml(filePath string, storage interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return parse(data, storage)
}

func merge(a types.Keywords, b *types.Keywords) types.Keywords {
	for k := range *b {
		a[k] = (*b)[k]
	}
	return a
}

// GetKeywords returns keyword list with levels
func GetKeywords() types.Keywords {
	var config types.Config
	keywords := defaultConfig

	for _, path := range ConfigPaths {
		if fileExists(path) {
			err := readYaml(path, &config)
			if err != nil {
				continue
			}
			keywords = merge(keywords, &config.Keywords)
		}
	}
	return keywords
}