package main

import (
	"os"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/parser"
	"github.com/mishamyrt/checode/v1/pkg/reporters"
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
var reportFormat string
var outputFileName string

func init() {
	pflag.StringVarP(&configPath, "config", "c", "", "config file path")
	pflag.StringVarP(&reportFormat, "report", "r", "", "report format")
	pflag.StringVarP(&outputFileName, "output", "o", "", "output file name")
	pflag.Parse()
}

func main() {
	var currentConfig types.Config = config.GetConfig(configPath)
	parsingResulut := parser.Parse(pflag.Args(), currentConfig)
	switch reportFormat {
	case "md":
		reporters.WriteReport(reportFormat, outputFileName, parsingResulut)
	}
	exit(!bit.IsSet(parsingResulut.Flags, config.ErrFlag))
}
