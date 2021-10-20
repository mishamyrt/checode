package warnings

type commandHandler struct {
	Keyword string
	Handle  func(m *Match, argument string, message string) string
}
