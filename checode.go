package main

import (
	"os"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/file"
	"github.com/mishamyrt/checode/v1/pkg/paths"
	"github.com/mishamyrt/checode/v1/pkg/stdout"
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
	parsing := file.Parsing{
		Config: config.GetConfig(configPath),
	}

	requestedPaths := pflag.Args()
	if len(requestedPaths) == 0 {
		requestedPaths = []string{"."}
	}

	files := paths.Collect(requestedPaths)
	parsing.Run(files)

	stdout.Print(&parsing)
	exit(!parsing.Flags.IsSet(config.ErrFlag))
}
