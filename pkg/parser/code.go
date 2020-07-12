package parser

import (
	"bufio"
	"os"
	"strings"
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func getSplited(s string) (r string) {
	index := strings.Index(s, ":")
	if index == -1 {
		return
	}
	return s[:index]
}

func matchKeywords(s string, config *types.Config, match *types.Match) {
	var index int
	var substring string
	var hasKeyword = false
	var part = getSplited(s)
	if len(part) == 0 && len(match.Keywords) == 0 {
		return
	}
	for keyword := range *config {
		index = strings.Index(part, keyword)
		if index == -1 || index != (len(part)-len(keyword)) {
			continue
		}
		hasKeyword = true
		substring = s[index+len(keyword)+1:]
		match.Keywords = append((*match).Keywords, keyword)
		matchKeywords(substring, config, match)
		break
	}
	if !hasKeyword && len(match.Keywords) > 0 {
		match.Message = s
	}
}

func getLevelMap(keywords []string, config *types.Config) (result uint8) {
	for _, kw := range keywords {
		result |= (*config)[kw]
	}
	return
}

func parseFile(path string, cg *types.Config, c chan types.FileMatches, wg *sync.WaitGroup) {
	var matches []types.Match
	var line = 0
	var flags uint8 = 0

	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		var match types.Match
		matchKeywords(scanner.Text(), cg, &match)
		if len(match.Keywords) > 0 {
			match.Line = line
			match.Flags = getLevelMap(match.Keywords, cg)
			flags |= match.Flags
			matches = append(matches, match)
		}
	}

	c <- types.FileMatches{
		Matches: matches,
		Flags:   flags,
		Path:    path,
	}
}
