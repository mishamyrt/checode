package parser

import (
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/configuration"
	"github.com/mishamyrt/checode/v1/pkg/paths"
	"github.com/mishamyrt/checode/v1/pkg/stdout"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// Parse given paths
func Parse(filePaths []string) bool {
	var wg sync.WaitGroup
	success := true
	processedCount := 0
	c := make(chan types.FileMatches)

	config := configuration.GetConfiguration()
	filePaths = paths.CollectPaths(filePaths)

	wg.Add(len(filePaths))
	for _, path := range filePaths {
		go ParseFile(path, &config, c, &wg)
	}

	for i := range c {
		processedCount++
		if processedCount == len(filePaths) {
			close(c)
		}
		if len(i.Matches) == 0 {
			continue
		}
		stdout.PrintMatch(i.Path, i.Matches)
		success = success && i.Success
		return success
	}
	wg.Wait()
	return success
}
