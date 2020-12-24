package substring

import "strings"

// Trim given comment
func Trim(text string) string {
	return strings.Trim(text, "\n\t *")
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
