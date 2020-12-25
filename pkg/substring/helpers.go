package substring

import "strings"

// Trim given comment
func Trim(text string) string {
	return strings.Trim(text, "\n\t *")
}

// GetMidst text
func GetMidst(startDel string, endDel string, text string) string {
	startOffset := strings.Index(text, startDel)
	endOffset := strings.Index(text, endDel)
	if startOffset < 0 || endOffset < 0 {
		return ""
	}
	return text[startOffset+len(startDel) : endOffset]
}

// GetSubsequent text
func GetSubsequent(delimeter string, text string) string {
	offset := strings.Index(text, delimeter)
	if offset < 0 {
		return ""
	}
	return text[offset+len(delimeter):]
}

// GetPrevious text
func GetPrevious(delimeter string, text string) string {
	offset := strings.Index(text, delimeter)
	if offset < 0 {
		return ""
	}
	return text[:offset]
}
