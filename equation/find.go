package equation

type Finder interface {
	Find(term string) string
}

func Find(term string, finder Finder, logger Logger) string {
	logger.Debug("Finding equation...")

	return finder.Find(term)
}
