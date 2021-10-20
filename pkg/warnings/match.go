package warnings

import (
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/substring"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

var commands = []commandHandler{
	DueHandler,
}

// Match is basic TODO, FIXME, etc message.
type Match struct {
	Keywords []string
	Flags    bit.Map
	Line     int
	Message  string
}

func (m *Match) processMessage(s string) {
	message := substring.Trim(s)
	for _, command := range commands {
		seq := ":" + command.Keyword + " "
		if !strings.HasPrefix(message, seq) {
			continue
		}
		seqLength := len(seq)
		colonIndex := strings.Index(message[seqLength:], ":")
		if colonIndex <= 0 {
			continue
		}
		argument := substring.Trim(message[seqLength-1 : colonIndex+seqLength])
		tail := substring.Trim(message[colonIndex+seqLength+1:])
		message = command.Handle(m, argument, tail)
		break
	}
	m.Message = message
}

// Parse keywords from given string.
func (m *Match) Parse(s string, config *types.Config) {
	index := strings.Index(s, ":")
	if index < 0 {
		m.processMessage(s)
		return
	}

	keyword, flag := findKeyword(config, substring.Trim(s[0:index]))
	if keyword == "" {
		m.processMessage(s)
		return
	}

	m.Keywords = append(m.Keywords, keyword)
	m.Flags |= bit.Map(flag)
	if len(s) > index+1 {
		m.Parse(s[index+1:], config)
	}
}
