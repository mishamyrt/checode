package main

import (
	"os"

	"github.com/mishamyrt/checode/v1/pkg/configuration"
	"github.com/mishamyrt/checode/v1/pkg/files"
	"github.com/mishamyrt/checode/v1/pkg/formatter"
	"github.com/mishamyrt/checode/v1/pkg/parser"
)

func exit(success bool) {
	var exitCode = 0
	if !success {
		exitCode = 1
	}
	os.Exit(exitCode)
}

func main() {
	config := configuration.GetConfiguration()
	var keywords = configuration.ExtractKeywords(&config)

	// TODO: Use directory from args
	paths, _ := files.GetFiles(".")

	success := true

	// TODO: Process files asynchronously
	for _, path := range paths {
		matches, err := parser.Parse(path, &keywords)
		if err != nil || len(matches) == 0 {
			continue
		}

		succeeded := formatter.PrintMatch(path, matches, config)
		success = success && succeeded
	}
	exit(success)
}
