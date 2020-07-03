package types

// STOPSHIP: Add real description

// Match lol
type Match struct {
	Keyword string
	Line    int
	Message string
}

// KeywordList map of keyword with error level
type KeywordList map[string]string

// Configuration is list with keywords and error levels
type Configuration struct {
	Keywords KeywordList
}
