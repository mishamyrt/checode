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

var Cases = []TestCase{{
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
}}

func scannerFrom(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}

func Suite(t *testing.T, parse CommentParser) {
	for _, c := range Cases {
		res := parse(scannerFrom(c.Text), c.Set)
		if len(res) != c.Count {
			t.Fail()
		}
		for _, comment := range res {
			if comment != "Comment" {
				t.Fail()
			}
		}
	}
}

func TestParse(t *testing.T) {
	Suite(t, comments.Parse)
}
