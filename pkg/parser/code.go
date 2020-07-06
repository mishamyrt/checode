package parser

import (
	"bufio"
	"os"
	"strings"
	"sync"

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

// ParseFile given code
func ParseFile(path string, keywordList *types.KeywordList, c chan types.FileMatches, wg *sync.WaitGroup) {
	var matches []types.Match
	success := true

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()
	defer wg.Done()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		match := matchKeyword(scanner.Text(), keywordList, line)
		if len(match.Keyword) > 0 {
			match.Level = (*keywordList)[match.Keyword]
			matches = append(matches, match)
			success = success && match.Level != "err"
		}
	}
	c <- types.FileMatches{
		Matches: matches,
		Success: success,
		Path:    path,
	}
}
