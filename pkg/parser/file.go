package parser

import (
	"bufio"
	"io"
	"os"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// FileMatches is the line file results
type FileMatches struct {
	Matches []LineMatch
	Path    string
	Flags   bit.Map
}

// Parse given reader content
func (m *FileMatches) Parse(file io.Reader, config *types.Config) {
	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		match := LineMatch{Line: line}
		match.Parse(scanner.Text(), config)
		if len(match.Keywords) > 0 {
			m.Flags |= match.Flags
			m.Matches = append(m.Matches, match)
		}
	}
}

// ParseFile parses given file
func ParseFile(path string, config *types.Config) FileMatches {
	matches := FileMatches{Path: path}

	file, err := os.Open(path)
	if err != nil {
		return matches
	}

	defer file.Close()
	matches.Parse(file, config)
	return matches
}
