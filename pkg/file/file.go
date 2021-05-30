package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/colors"
	"github.com/mishamyrt/checode/v1/pkg/types"
	"github.com/mishamyrt/checode/v1/pkg/warnings"

	"github.com/mishamyrt/compars"
	"github.com/mishamyrt/compars/pkg/symbols"
	ctypes "github.com/mishamyrt/compars/pkg/types"
)

// Matches is the line file results.
type Matches struct {
	Matches []warnings.Match
	Path    string
	Flags   bit.Map
}

// Parse given reader content.
func Parse(file io.Reader, config *types.Config, set ctypes.CommentSymbolSet) (m Matches) {
	scanner := bufio.NewScanner(file)
	r := compars.Parse(scanner, set)
	for _, s := range r {
		var match warnings.Match
		if s.Text == "" {
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

// ParseFile parses given file.
func ParseFile(path string, config *types.Config) (matches Matches) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	ext := filepath.Ext(path)
	if ext == "" {
		return
	}
	set, err := symbols.GetSetByExtension(ext)
	if err != nil {
		fmt.Println(colors.Yellow("Unknown extension: " + ext))
		return
	}
	matches = Parse(file, config, set)
	matches.Path = path
	return
}
