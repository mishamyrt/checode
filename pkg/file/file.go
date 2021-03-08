package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/comments"
	"github.com/mishamyrt/checode/v1/pkg/types"
	"github.com/mishamyrt/checode/v1/pkg/warnings"
)

// Matches is the line file results
type Matches struct {
	Matches []warnings.Match
	Path    string
	Flags   bit.Map
}

// Parse given reader content
func Parse(file io.Reader, config *types.Config, set types.CommentSymbolSet) (m Matches) {
	scanner := bufio.NewScanner(file)
	r := comments.Parse(scanner, set)
	for _, s := range r {
		var match warnings.Match
		if len(s.Text) == 0 {
			continue
		}
		match.Parse(s.Text, config)
		if len(match.Keywords) == 0 {
			continue
		}
		match.Line = s.Line
		m.Flags |= match.Flags
		m.Matches = append(m.Matches, match)
	}
	return
}

// ParseFile parses given file
func ParseFile(path string, config *types.Config) (matches Matches) {
	// fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return
	}
	set, err := getSetByExtension(filepath.Ext(path))
	if err != nil {
		return
	}
	matches = Parse(file, config, set)
	matches.Path = path
	defer file.Close()
	return
}
