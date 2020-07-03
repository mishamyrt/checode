package main

import (
	"os"

	"github.com/mishamyrt/checode/v1/pkg/configuration"
	"github.com/mishamyrt/checode/v1/pkg/files"
	"github.com/mishamyrt/checode/v1/pkg/formatter"
	"github.com/mishamyrt/checode/v1/pkg/parser"
)

func exit(success bool) {
	exitCode := 0
	if !success {
		exitCode = 1
	}
	os.Exit(exitCode)
}

func main() {
	var paths []string
	success := true
	requestedPaths := os.Args[1:]

	config := configuration.GetConfiguration()
	keywords := configuration.ExtractKeywords(&config)

	if len(requestedPaths) == 0 {
		requestedPaths = append(requestedPaths, ".")
	}

	for _, path := range requestedPaths {
		files, _ := files.GetFiles(path)
		if len(files) > 0 {
			paths = append(paths, files...)
		}
	}

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
