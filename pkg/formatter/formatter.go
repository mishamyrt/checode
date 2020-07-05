package formatter

import (
	"fmt"
	"strconv"

	"github.com/mishamyrt/checode/v1/pkg/colours"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// PrintMatch to stdout
func PrintMatch(path string, matches []types.Match, rules types.KeywordList) bool {
	var kw string
	result := colours.Underline(fmt.Sprintln(path))
	success := true

	for _, match := range matches {
		result += colours.Grey(fmt.Sprintf("  %-4s", strconv.Itoa(match.Line)))
		switch rules[match.Keyword] {
		case "err":
			kw = colours.Red(match.Keyword)
			success = false
		case "warn":
			kw = colours.Yellow(match.Keyword)
		default:
			kw = colours.Blue(match.Keyword)
		}
		result += fmt.Sprintf("%18s", kw)
		result += fmt.Sprintf("  %s", match.Message)
		result += "\n"
	}
	fmt.Println(result)
	// TODO: Move status handling outside formatter
	return success
}
