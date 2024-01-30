package equation

type Logger interface {
	Debug(description string)
	Error(description string)
	Info(description string)
	Warn(description string)
}

func Description() string {
	return "Package provides operations around math equations"
}
