package config

import (
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// DefaultConfigPath is allowed configuration file names
const DefaultConfigPath = ".checode.yaml"

// DefaultConfig is a built-in basic configuration
var DefaultConfig = types.Config{
	"FIXME":    types.ErrKeyword,
	"STOPSHIP": types.ErrKeyword,
	"TODO":     types.WarnKeyword,
}

// GetConfig Reads the custom configuration file, then the file along the standard path.
// If nothing is found, it returns the standard value.
func GetConfig(userConfigPath string) types.Config {
	var result types.Config
	var configPaths = []string{DefaultConfigPath, userConfigPath}
	var err error

	for _, path := range configPaths {
		if len(path) == 0 {
			continue
		}
		err = ReadConfigFile(path, &result)
		if err == nil {
			return result
		}
	}
	return DefaultConfig
}
