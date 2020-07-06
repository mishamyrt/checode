package main

import (
	"os"

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
	exit(parser.Parse(os.Args[1:]))
}
