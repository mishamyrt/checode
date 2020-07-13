package types

// FileMatches is file parsing result
type FileMatches struct {
	Matches []Match
	Path    string
	Flags   uint8
}

// Match result
type Match struct {
	Keywords []string
	Flags    uint8
	Line     int
	Message  string
}

// Config map of keyword with error level
type Config map[string]uint8

// KewordsMap is a map where key is keyword and value is level
type KewordsMap map[string]string

// ParsingResult for reporters
type ParsingResult struct {
	FileMatches []FileMatches
	Flags       uint8
}

// ConfigFile is list with keywords and error levels
type ConfigFile struct {
	Keywords KewordsMap
}
