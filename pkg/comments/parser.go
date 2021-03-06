package comments

import (
	"bufio"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/substring"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// Parse comments from given scanner
func Parse(s *bufio.Scanner, set types.CommentSymbolSet) []types.FileComments {
	var results []types.FileComments
	lineNumber := 0
	multilineStart := 0
	line := ""
	comment := ""
	inMultiline := false

	hasInline := len(set.Inline) > 0

	for s.Scan() {
		lineNumber++
		line = s.Text()

		if !inMultiline && strings.Contains(line, set.MultilineStart) {
			subLine := substring.GetSubsequent(set.MultilineStart, line)
			if strings.Contains(subLine, set.MultilineEnd) {
				results = append(results, types.FileComments{
					Text: substring.Trim(
						substring.GetPrevious(set.MultilineEnd, subLine)),
					Line: lineNumber,
				})
			} else {
				inMultiline = true
				multilineStart = lineNumber
				comment += substring.GetSubsequent(set.MultilineStart, line) + "\n"
			}
			continue
		}

		if inMultiline && strings.Contains(line, set.MultilineEnd) {
			inMultiline = false
			comment += substring.GetPrevious(set.MultilineEnd, line)
			results = append(results, types.FileComments{
				Text: substring.Trim(comment),
				Line: multilineStart,
			})
			comment = ""
			continue
		} else if inMultiline {
			comment += line + "\n"
			continue
		}

		if hasInline && strings.Contains(line, set.Inline) {
			results = append(results,
				types.FileComments{
					Text: substring.Trim(
						substring.GetSubsequent(set.Inline, line)),
					Line: lineNumber,
				})
		}
	}
	return results
}
