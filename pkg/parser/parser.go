package parser

import (
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/paths"
	"github.com/mishamyrt/checode/v1/pkg/stdout"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// Parse given paths
func Parse(filePaths []string) bool {
	var wg sync.WaitGroup
	var processedCount = 0
	var success = true
	var matchesChan = make(chan types.FileMatches)

	keywords := config.GetKeywords()
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

	for i := range matchesChan {
		processedCount++
		if processedCount == len(filePaths) {
			close(matchesChan)
		}
		if len(i.Matches) == 0 {
			continue
		}
		stdout.PrintMatch(i.Path, i.Matches)
		success = success && i.Success
	}
	wg.Wait()
	return success
}
