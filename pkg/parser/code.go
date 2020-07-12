package parser

import (
	"bufio"
	"os"
	"strings"
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func matchKeyword(s string, keywords *types.Config, line int) types.Match {
	var index int
	for keyword := range *keywords {
		index = strings.Index(s, keyword+":")
		if index > -1 {
			return types.Match{
				Keyword: keyword,
				Line:    line,
				Message: strings.Trim(s[index+len(keyword)+1:], " "),
			}
		}
	}
	return types.Match{}
}

func parseFile(path string, keywords *types.Config, c chan types.FileMatches, wg *sync.WaitGroup) {
	var matches []types.Match
	var success = true
	var line = 0

	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		match := matchKeyword(scanner.Text(), keywords, line)
		if len(match.Keyword) > 0 {
			match.Level = (*keywords)[match.Keyword]
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
