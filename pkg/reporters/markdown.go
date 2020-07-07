package reporters

import (
	"strconv"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

// FormatMarkdown returns markdown formatted matches
func FormatMarkdown(m types.FileMatches) (result string) {
	result += "__" + m.Path + "\n" + "__"
	result += "Line |Type |Message \n"
	result += "---|---|---\n"
	for _, match := range m.Matches {
		result += strconv.Itoa(match.Line) + "| "
		result += match.Keyword + "| "
		result += match.Message + "\n"
	}
	return
}
