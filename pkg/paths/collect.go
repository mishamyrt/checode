package paths

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/config"
)

func isAllowedPath(path string) bool {
	for _, v := range config.ConfigPaths {
		if strings.HasSuffix(path, v) {
			return false
		}
	}
	return !strings.Contains(path, ".git/")
}

func getFiles(path string) (paths []string, err error) {
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

// CollectPaths of files
func CollectPaths(paths []string) (result []string) {
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	for _, path := range paths {
		files, _ := getFiles(path)
		if len(files) > 0 {
			result = append(result, files...)
		}
	}
	return
}
