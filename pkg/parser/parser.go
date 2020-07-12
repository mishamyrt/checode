package parser

import (
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/paths"
	"github.com/mishamyrt/checode/v1/pkg/reporters"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// Parse given paths
func Parse(filePaths []string, keywords types.Config) bool {
	var wg sync.WaitGroup
	var processedCount = 0
	var success = true
	var matchesChan = make(chan types.FileMatches)

	filePaths = paths.CollectPaths(filePaths)

	// Early exit if none
	if len(filePaths) == 0 {
		close(matchesChan)
		return true
	}

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
		reporters.PrintMatch(r)
		success = success && r.Success
	}
	wg.Wait()
	return success
}
