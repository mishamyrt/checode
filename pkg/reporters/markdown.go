package reporters

import (
	"strconv"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/parser"
)

// FormatMarkdown returns markdown formatted matches
func FormatMarkdown(data *parser.Parsing) (result string) {
	for _, file := range data.Matches {
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
