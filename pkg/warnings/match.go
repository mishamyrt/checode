package warnings

import (
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/substring"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

func findKeyword(a *types.Config, x string) (keyword string, flag uint8) {
	for i, m := range *a {
		if x == i {
			return i, m
		}
	}
	return "", 0
}

// Match is basic TODO, FIXME, etc message.
type Match struct {
	Keywords []string
	Flags    bit.Map
	Line     int
	Message  string
}

// Parse keywords from given string.
func (m *Match) Parse(s string, config *types.Config) {
	index := strings.Index(s, ":")
	if index < 0 {
		m.Message = substring.Trim(s)
		return
	}

	keyword, flag := findKeyword(config, substring.Trim(s[0:index]))
	if keyword == "" {
		m.Message = substring.Trim(s)
		return
	}

	m.Keywords = append(m.Keywords, keyword)
	m.Flags |= bit.Map(flag)
	if len(s) > index+1 {
		m.Parse(s[index+1:], config)
	}
}
