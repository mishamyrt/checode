package reporters

import (
	"strconv"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

// FormatMarkdown returns markdown formatted matches
func FormatMarkdown(data types.ParsingResult) (result string) {
	for _, file := range data.FileMatches {
		result += "__" + file.Path + "__" + "\n"
		result += "Line |Type |Message \n"
		result += "---|---|---\n"
		for _, match := range file.Matches {
			result += strconv.Itoa(match.Line) + "| "
			result += strings.Join(match.Keywords, " ") + "| "
			result += match.Message + "\n"
		}
		result += "\n"
	}
	return
}
