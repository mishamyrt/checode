package configuration

import (
	"io/ioutil"
	"os"

	"github.com/mishamyrt/checode/v1/pkg/types"
	"gopkg.in/yaml.v2"
)

var configurationPath = []string{".checode.yml", ".checode.yaml"}

var defaultConfiguration = types.KeywordList{
	"FIXME":    "err",
	"STOPSHIP": "err",
	"TODO":     "warn",
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

func merge(a types.KeywordList, b types.KeywordList) types.KeywordList {
	for k := range b {
		a[k] = b[k]
	}
	return a
}

// ExtractKeywords from given list
func ExtractKeywords(keywordList *types.KeywordList) []string {
	keys := make([]string, 0, len(*keywordList))
	for k := range *keywordList {
		keys = append(keys, k)
	}
	return keys
}

// GetConfiguration returns keyword list with levels
func GetConfiguration() types.KeywordList {
	var config types.Configuration
	var keywords = defaultConfiguration

	for _, path := range configurationPath {
		if fileExists(path) {
			err := readYaml(path, &config)
			if err != nil {
				continue
			}
			keywords = merge(keywords, config.Keywords)
		}
	}
	return defaultConfiguration
}
