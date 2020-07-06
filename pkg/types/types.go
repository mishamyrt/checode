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

// Keywords map of keyword with error level
type Keywords map[string]string

// Config is list with keywords and error levels
type Config struct {
	Keywords Keywords
}
