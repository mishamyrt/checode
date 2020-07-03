package config

import (
	"io/ioutil"

	"github.com/mishamyrt/checode/v1/pkg/types"
	"gopkg.in/yaml.v2"
)

// STOPSHIP: Add predefined configuration

// ExtractKeywords from given list
func ExtractKeywords(keywordList *types.KeywordList) []string {
	keys := make([]string, 0, len(*keywordList))
	for k := range *keywordList {
		keys = append(keys, k)
	}
	return keys
}

// Parse YAML raw bytes data to given link
func Parse(data []byte, storage interface{}) error {
	return yaml.Unmarshal(data, storage)
}

// ReadFile reads YAML file to given link
func ReadFile(filePath string, storage interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return Parse(data, storage)
}
