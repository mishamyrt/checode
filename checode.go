package main

import (
	"os"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/parser"
	"github.com/mishamyrt/checode/v1/pkg/types"
	"github.com/spf13/pflag"
)

func exit(success bool) {
	exitCode := 0
	if !success {
		exitCode = 1
	}
	os.Exit(exitCode)
}

var configPath string

func init() {
	pflag.StringVarP(&configPath, "config", "c", "", "config file path")
	pflag.Parse()
}

func main() {
	var currentConfig types.Config = config.GetConfig(configPath)

	exit(parser.Parse(pflag.Args(), currentConfig))
}
