package comments

import "github.com/mishamyrt/checode/v1/pkg/types"

// CommentSymbols is allowed comment symbol set
var CommentSymbols = map[string]types.CommentSymbolSet{
	"c-style": {
		Inline:         "//",
		MultilineStart: "/*",
		MultilineEnd:   "*/",
	},
	"python": {
		Inline:         "#",
		MultilineStart: "'''",
		MultilineEnd:   "'''",
	},
	"html": {
		Inline:         "",
		MultilineStart: "<!--",
		MultilineEnd:   "-->",
	},
}
