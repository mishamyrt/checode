package parser

import (
	"bufio"
	"os"
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func matchKeyword(s string, kl *types.KeywordList, line int) (match types.Match) {
	var index int
	for k := range *kl {
		index = strings.Index(s, k+":")
		if index > -1 {
			match.Keyword = k
			match.Line = line
			match.Message = s[index+len(k)+1:]
			return
		}
	}
	return
}

// Parse given code
func Parse(path string, keywordList *types.KeywordList) (matches []types.Match, success bool) {
	success = true

	file, err := os.Open(path)
	if err != nil {
		return nil, false
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		match := matchKeyword(scanner.Text(), keywordList, line)
		if len(match.Keyword) > 0 {
			matches = append(matches, match)
			success = success && (*keywordList)[match.Keyword] != "err"
		}
	}
	return
}
