package file_test

import (
	"io"
	"testing"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/file"
	"github.com/mishamyrt/checode/v1/pkg/testhelpers"
	"github.com/mishamyrt/checode/v1/pkg/types"
	"github.com/mishamyrt/compars/pkg/symbols"
)

type Parser func(file io.Reader, config *types.Config) file.Matches

func TestParse(t *testing.T) {
	testComment := "# NOTE: Test"
	set, _ := symbols.GetSetByExtension(".sh")
	result := file.Parse(testhelpers.ReaderFrom(testComment), &config.DefaultConfig, set)
	if len(result.Matches) != 1 {
		t.Fail()
	}
}
