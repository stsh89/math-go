package equation

type Deleter interface {
	Delete(term string) string
}

func Delete(term string, deleter Deleter, logger Logger) string {
	logger.Info("Deleting equation...")

	return deleter.Delete(term)
}
