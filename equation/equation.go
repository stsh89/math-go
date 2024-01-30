package equation

type Logger interface {
	Debug(description string, args ...any)
	Error(description string, args ...any)
	Info(description string, args ...any)
	Warn(description string, args ...any)
}

func Description() string {
	return "Package provides operations around math equations"
}
