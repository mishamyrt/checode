package formatter

import (
	"fmt"
	"strconv"

	"github.com/mishamyrt/checode/v1/pkg/colors"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// PrintMatch to stdout
func PrintMatch(path string, matches []types.Match, rules types.KeywordList) bool {
	var result = colors.Underline(fmt.Sprintln(path))
	success := true
	var kw string
	for _, match := range matches {
		result += colors.Grey(fmt.Sprintf("  %-4s", strconv.Itoa(match.Line)))
		switch rules[match.Keyword] {
		case "err":
			kw = colors.Red(match.Keyword)
			success = false
		case "warn":
			kw = colors.Yellow(match.Keyword)
		default:
			kw = colors.Blue(match.Keyword)
		}
		result += fmt.Sprintf("%18s", kw)
		result += fmt.Sprintf("  %s", match.Message)
		result += "\n"
	}
	fmt.Println(result)
	// TODO: Move status handling outside formatter
	return success
}
