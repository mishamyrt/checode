package main

import (
	"fmt"
	"os"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/parser"
	"github.com/mishamyrt/checode/v1/pkg/paths"
	"github.com/mishamyrt/checode/v1/pkg/reporters"
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
var reportFormat string
var outputFileName string

func init() {
	pflag.StringVarP(&configPath, "config", "c", "", "config file path")
	pflag.StringVarP(&reportFormat, "report", "r", "", "report format")
	pflag.StringVarP(&outputFileName, "output", "o", "", "output file name")
	pflag.Parse()
}

func main() {
	parsingResulut := parser.Parsing{
		Config: config.GetConfig(configPath),
	}
	parsingResulut.Run(paths.CollectPaths(pflag.Args()))
	stdout.PrintMatches(&parsingResulut)
	if len(reportFormat) > 0 {
		err := reporters.CreateReport(reportFormat, outputFileName, &parsingResulut)
		if err != nil {
			fmt.Println(err.Error())
			exit(false)
		}
	}
	exit(!bit.IsSet(parsingResulut.Flags, config.ErrFlag))
}
