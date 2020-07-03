package files

import (
	"os"
	"path/filepath"
	"strings"
)

// GetFiles collects files from given folder recursively
func GetFiles(path string) (paths []string, err error) {
	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// FIXME: Use configuration name from list
			if !info.IsDir() && !strings.HasSuffix(path, ".checode.yaml") {
				paths = append(paths, path)
			}
			return nil
		})
	return
}
