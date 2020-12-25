package comments

import (
	"bufio"
	"fmt"
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

	hasInline := len(set.Inline) > 0

	for s.Scan() {
		lineNumber++
		line = s.Text()

		if !inMultiline && contains(line, set.MultilineStart) {
			subLine := substring.GetSubsequent(set.MultilineStart, line)
			if contains(subLine, set.MultilineEnd) {
				results = append(results, substring.Trim(
					substring.GetPrevious(set.MultilineEnd, subLine)))
			} else {
				inMultiline = true
				comment += substring.GetSubsequent(set.MultilineStart, line)
			}
			continue
		}

		if inMultiline && contains(line, set.MultilineEnd) {
			inMultiline = false
			comment += substring.GetPrevious(set.MultilineEnd, line)
			results = append(results, substring.Trim(comment))
			comment = ""
			continue
		} else if inMultiline {
			comment += line
			continue
		}

		if hasInline && contains(line, set.Inline) {
			results = append(results,
				substring.Trim(
					substring.GetSubsequent(set.Inline, line)))
		}
	}
	fmt.Println(results)
	return results
}
