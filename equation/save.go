package equation

type Saver interface {
	Save(term string) string
}

func Save(term string, saver Saver, logger Logger) string {
	logger.Info("Saving equation...")

	return saver.Save(term)
}
