package parser

import (
	"strings"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// LineMatch is the line parsing results
type LineMatch struct {
	Keywords []string
	Flags    bit.Map
	Line     int
	Message  string
}

// Parse line with given config
func (m *LineMatch) Parse(text string, config *types.Config) {
	colonIndex := strings.Index(text, ":")
	// NOTE: Exit if colon not found
	if colonIndex < 0 {
		// NOTE: Save message if have keywords
		if len(m.Keywords) > 0 {
			m.Message = text
		}
		return
	}

	lineSlice := text[:colonIndex]
	keywordIndex := 0

	for keyword := range *config {
		keywordIndex = strings.Index(lineSlice, keyword)

		// NOTE: Skip if keyword not found
		expectedOffset := len(lineSlice) - len(keyword)
		if keywordIndex == -1 || keywordIndex != expectedOffset {
			continue
		}
		m.Keywords = append(m.Keywords, keyword)
		m.Flags |= bit.Map((*config)[keyword])

		// NOTE: Rerun function with stripped line
		m.Parse(text[keywordIndex+len(keyword)+1:], config)
		break
	}
}
