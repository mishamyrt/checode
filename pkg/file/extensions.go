package file

import (
	"errors"

	"github.com/mishamyrt/checode/v1/pkg/comments"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// ExtensionsCommentSet is map with file extensions and comment sets
var ExtensionsCommentSet = map[string]types.CommentSymbolSet{
	"js":   comments.CommentSymbols["c-style"],
	"jsx":  comments.CommentSymbols["c-style"],
	"ts":   comments.CommentSymbols["c-style"],
	"tsx":  comments.CommentSymbols["c-style"],
	"go":   comments.CommentSymbols["c-style"],
	"py":   comments.CommentSymbols["python"],
	"html": comments.CommentSymbols["html"],
}

func getSetByExtension(ext string) (types.CommentSymbolSet, error) {
	if val, ok := ExtensionsCommentSet[ext[1:]]; ok {
		return val, nil
	}
	return types.CommentSymbolSet{}, errors.New("Value not found")
}
