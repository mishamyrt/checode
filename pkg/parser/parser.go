package parser

import (
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/paths"
	"github.com/mishamyrt/checode/v1/pkg/stdout"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// Parse given paths
func Parse(filePaths []string, keywords types.Config) types.ParsingResult {
	var wg sync.WaitGroup
	var processedCount = 0
	var result types.ParsingResult
	var matchesChan = make(chan types.FileMatches)

	filePaths = paths.CollectPaths(filePaths)

	// Early exit if none
	if len(filePaths) == 0 {
		close(matchesChan)
		return result
	}

	// TODO: asd

	// Fill work group with paths length
	wg.Add(len(filePaths))

	for _, path := range filePaths {
		go parseFile(path, &keywords, matchesChan, &wg)
	}

	for r := range matchesChan {
		processedCount++
		if processedCount == len(filePaths) {
			close(matchesChan)
		}
		if len(r.Matches) == 0 {
			continue
		}
		result.FileMatches = append(result.FileMatches, r)
		stdout.PrintMatch(r)
		result.Flags |= r.Flags
	}
	wg.Wait()
	return result
}
