package comments

import (
	"bufio"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/substring"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

func contains(haystack string, needle string) bool {
	return strings.Index(haystack, needle) >= 0
}

// Parse comments from given scanner
func Parse(s *bufio.Scanner, set types.CommentSymbolSet) []string {
	var results []string
	lineNumber := 0
	line := ""
	comment := ""
	inMultiline := false

	for s.Scan() {
		lineNumber++
		line = s.Text()

		if contains(line, set.MultilineEnd) {
			inMultiline = false
			comment += substring.GetSubsequent(set.MultilineEnd, line)
			results = append(results, substring.Trim(comment))
			comment = ""
			continue
		} else if inMultiline {
			comment += line
			continue
		}

		if contains(line, set.MultilineStart) {
			inMultiline = true
			comment += substring.GetSubsequent(set.MultilineStart, line)
			continue
		}

		if contains(line, set.Inline) {
			results = append(results,
				substring.Trim(
					substring.GetSubsequent(set.Inline, line)))
		}
	}
	return results
}
