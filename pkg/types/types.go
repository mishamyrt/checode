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

type KewordsMap map[string]string

// ConfigFile is list with keywords and error levels
type ConfigFile struct {
	Keywords KewordsMap
}
