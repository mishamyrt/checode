package parser

import (
	"bufio"
	"os"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func matchKeyword(s *string, kk *[]string, line int) (match types.Match) {
	var index int
	for _, k := range *kk {
		index = strings.Index(*s, k+":")
		if index > -1 {
			match.Keyword = k
			match.Line = line
			match.Message = (*s)[index+len(k)+1:]
			return
		}
	}
	return
}

// Parse given code
func Parse(path string, keywords *[]string) (matches []types.Match, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	var text string
	for scanner.Scan() {
		line++
		text = scanner.Text()
		match := matchKeyword(&text, keywords, line)
		if len(match.Message) > 0 {
			matches = append(matches, match)
		}
	}
	err = scanner.Err()
	return
}
