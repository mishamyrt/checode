package stdout

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func printLine(n int) string {
	return grey(fmt.Sprintf("  %-5s", strconv.Itoa(n)))
}

func printPath(p string) string {
	return underline(fmt.Sprintln(p))
}

func printKeywords(k []string, bitmap uint8) (kw string) {
	kw = strings.Join(k, ": ") + ":"
	kw = colorize(bitmap)(kw)
	return
}

func printMessage(m string) string {
	return fmt.Sprintf(" %s", m)
}

// PrintMatch to stdout
func PrintMatch(m types.FileMatches) {
	result := printPath(m.Path)
	for _, match := range m.Matches {
		result += printLine(match.Line)
		result += printKeywords(match.Keywords, match.Flags)
		result += printMessage(match.Message)
		result += "\n"
	}
	fmt.Printf("%s\n", result)
}
