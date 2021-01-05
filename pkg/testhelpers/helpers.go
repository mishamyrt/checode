package testhelpers

import (
	"bufio"
	"strings"
)

// ReaderFrom creates strings.Reader instance from string
func ReaderFrom(input string) *strings.Reader {
	return strings.NewReader(input)
}

// ScannerFrom creates bufio.Scanner instance from string
func ScannerFrom(input string) *bufio.Scanner {
	return bufio.NewScanner(ReaderFrom(input))
}
