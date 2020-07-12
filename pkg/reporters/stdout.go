package reporters

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

func line(n int) string {
	return grey(fmt.Sprintf("  %-5s", strconv.Itoa(n)))
}

func path(p string) string {
	return underline(fmt.Sprintln(p))
}

func isSet(bitmap uint8, flag uint8) bool {
	return (bitmap & flag) == flag
}

func colorize(bitmap uint8) func(s string) string {
	if isSet(bitmap, config.ErrFlag) {
		return red
	} else if isSet(bitmap, config.WarnFlag) {
		return yellow
	}
	return blue
}

func keywords(k []string, bitmap uint8) string {
	var kw string = strings.Join(k, ": ") + ":"
	kw = colorize(bitmap)(kw)
	return fmt.Sprintf("%20s", kw)
}

func message(m string) string {
	return fmt.Sprintf(" %s", m)
}

// PrintMatch to stdout
func PrintMatch(m types.FileMatches) {
	result := path(m.Path)
	for _, match := range m.Matches {
		result += line(match.Line)
		result += keywords(match.Keywords, match.Flags)
		result += message(match.Message)
		result += "\n"
	}
	fmt.Printf("%s\n", result)
}
