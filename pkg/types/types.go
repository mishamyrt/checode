package types

const (
	// ErrKeyword key
	ErrKeyword = "err"
	// WarnKeyword key
	WarnKeyword = "warn"
)

// FileMatches is file parsing result
type FileMatches struct {
	Matches []Match
	Path    string
	Success bool
}

// Match result
type Match struct {
	Keyword string
	Level   string
	Line    int
	Message string
}

// Config map of keyword with error level
type Config map[string]string

// ConfigFile is list with keywords and error levels
type ConfigFile struct {
	Keywords Config
}
