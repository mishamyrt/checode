package parser

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/configuration"
)

func isAllowedPath(path string) bool {
	for _, v := range configuration.ConfigurationPath {
		if strings.HasSuffix(path, v) {
			return false
		}
	}
	return !strings.Contains(path, ".git/")
}

// GetFiles collects files from given folder recursively
func GetFiles(path string) (paths []string, err error) {
	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && isAllowedPath(path) {
				paths = append(paths, path)
			}
			return nil
		})
	return
}
