package types

// Config map of keyword with error level
type Config map[string]uint8

// KewordsMap is a map where key is keyword and value is level
type KewordsMap map[string]string

// ConfigFile is list with keywords and error levels
type ConfigFile struct {
	Keywords KewordsMap
}
