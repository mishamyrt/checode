package paths

import (
	"os"
	"path/filepath"
	"strings"
)

func isAllowedPath(path string) bool {
	if strings.HasSuffix(path, ".yaml") {
		return false
	} else if strings.HasSuffix(path, ".md") {
		return false
	}
	return !strings.Contains(path, ".git/")
}

// FileList is the list of files
type FileList []string

func (f *FileList) collectPath(path string) error {
	return filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && isAllowedPath(path) {
				*f = append(*f, path)
			}
			return nil
		})
}

// Collect files from given paths.
func Collect(paths []string) (list FileList) {
	for _, path := range paths {
		_ = list.collectPath(path)
	}
	return
}
