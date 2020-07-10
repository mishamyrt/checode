package reporters

import (
	"fmt"
	"strconv"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func line(n int) string {
	return grey(fmt.Sprintf("  %-4s", strconv.Itoa(n)))
}

func path(p string) string {
	return underline(fmt.Sprintln(p))
}

func keyword(k string, level string) string {
	var kw string
	switch level {
	case types.ErrKeyword:
		kw = red(k)
	case types.WarnKeyword:
		kw = yellow(k)
	default:
		kw = blue(k)
	}
	return fmt.Sprintf("%18s", kw)
}

func message(m string) string {
	return fmt.Sprintf("  %s", m)
}

// PrintMatch to stdout
func PrintMatch(m types.FileMatches) {
	result := path(m.Path)
	for _, match := range m.Matches {
		result += line(match.Line)
		result += keyword(match.Keyword, match.Level)
		result += message(match.Message)
		result += "\n"
	}
	fmt.Printf("%s\n", result)
}
