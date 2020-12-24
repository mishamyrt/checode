package comments

import (
	"bufio"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func trim(text string) string {
	return strings.Trim(text, "\n\t *")
}

func getSubsequent(delimeter string, text string) string {
	offset := strings.Index(text, delimeter)
	if offset < 0 {
		return ""
	}
	return text[offset+len(delimeter):]
}

func getPrevious(delimeter string, text string) string {
	offset := strings.Index(text, delimeter)
	if offset < 0 {
		return ""
	}
	return text[:offset]
}

func includes(haystack string, needle string) bool {
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

		if includes(line, set.MultilineEnd) {
			inMultiline = false
			comment += getPrevious(set.MultilineEnd, line)
			results = append(results, trim(comment))
			comment = ""
			continue
		} else if inMultiline {
			comment += line
			continue
		}

		if includes(line, set.MultilineStart) {
			inMultiline = true
			comment += getSubsequent(set.MultilineStart, line)
			continue
		}

		if includes(line, set.Inline) {
			results = append(results, trim(getSubsequent(set.Inline, line)))
		}
	}
	return results
}
