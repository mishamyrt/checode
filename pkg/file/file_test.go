package file_test

import (
	"io"
	"testing"

	"github.com/mishamyrt/checode/v1/pkg/comments"
	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/file"
	"github.com/mishamyrt/checode/v1/pkg/testhelpers"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

type Parser func(file io.Reader, config *types.Config) file.Matches

func TestParse(t *testing.T) {
	testComment := "// NOTE: Test"
	result := file.Parse(testhelpers.ReaderFrom(testComment), &config.DefaultConfig, comments.CommentSymbols["c-style"])
	if len(result.Matches) != 1 {
		t.Fail()
	}
}
