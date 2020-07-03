package main

import (
	"os"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/files"
	"github.com/mishamyrt/checode/v1/pkg/formatter"
	"github.com/mishamyrt/checode/v1/pkg/parser"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

func exit(success bool) {
	var exitCode = 0
	if !success {
		exitCode = 1
	}
	os.Exit(exitCode)
}

func main() {
	var configuration types.Configuration

	// TODO: Use configuration name from list
	err := config.ReadFile(".checode.yaml", &configuration)
	var keywords = config.ExtractKeywords(&configuration.Keywords)

	// FIXME: Add real error handling
	if err != nil {
		panic(err)
	}

	// FIXME: Add error handling
	// TODO: Use directory from args
	paths, _ := files.GetFiles(".")

	success := true

	// TODO: Process files asynchronously
	for _, path := range paths {
		matches, err := parser.Parse(path, &keywords)
		if err != nil || len(matches) == 0 {
			continue
		}

		succeeded := formatter.PrintMatch(path, matches, configuration.Keywords)
		success = success && succeeded
	}
	exit(success)
}
