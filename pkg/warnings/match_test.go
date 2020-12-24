package warnings_test

import (
	"strings"
	"testing"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/warnings"
)

func includes(a []string, x string) bool {
	for _, m := range a {
		if x == m {
			return true
		}
	}
	return false
}

func TestMatch(t *testing.T) {
	var m warnings.Match
	m.Parse("NOTE: FIXME: Text", &config.DefaultConfig)
	if m.Message != "Text" {
		t.Errorf("Wrong message: \"%s\"", m.Message)
		t.Fail()
	}
	if !includes(m.Keywords, "NOTE") || !includes(m.Keywords, "FIXME") {
		t.Errorf("Keyword not found: \"%s\"", strings.Join(m.Keywords, ", "))
		t.Fail()
	}
	if !m.Flags.IsSet(config.WarnFlag) {
		t.Errorf("Warn flag not found: \"%d\"", m.Flags)
		t.Fail()
	}
}
