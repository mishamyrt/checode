package bit

// Map is bitmap abstraction.
type Map uint8

// IsSet checks flag over value.
func (m *Map) IsSet(flag uint8) bool {
	return (uint8(*m) & flag) == flag
}
