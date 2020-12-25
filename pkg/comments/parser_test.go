package comments_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/mishamyrt/checode/v1/pkg/comments"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

type CommentParser func(s *bufio.Scanner, set types.CommentSymbolSet) []string

type TestCase struct {
	Count int
	Text  string
	Set   types.CommentSymbolSet
}

var Cases = []TestCase{
	{
		// For the first case, I set a symbol set directly to test the algorithm itself, not the data
		Set: types.CommentSymbolSet{
			Inline:         "//",
			MultilineStart: "/*",
			MultilineEnd:   "*/",
		},
		Count: 3,
		Text: `
			// Comment
			/*
			* Comment
			*/
			/*
			Comment
			*/`,
	},
	{
		Set:   comments.CommentSymbols["python"],
		Count: 2,
		Text: `
			# Comment
			'''
			Comment
			'''`,
	},
	{
		Set:   comments.CommentSymbols["html"],
		Count: 1,
		Text:  "<!-- Comment -->",
	},
}

func scannerFrom(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}

func ParseSuite(t *testing.T, parse CommentParser) {
	for _, c := range Cases {
		res := parse(scannerFrom(c.Text), c.Set)
		if len(res) != c.Count {
			t.Errorf("Wrong count: %d vs %d", len(res), c.Count)
			t.Fail()
		}
		for _, comment := range res {
			if comment != "Comment" {
				t.Errorf("Wrong text: \"%s\" from \"%s\"", comment, c.Text)
				t.Fail()
			}
		}
	}
}

func TestParse(t *testing.T) {
	ParseSuite(t, comments.Parse)
}
