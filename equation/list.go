package equation

type Lister interface {
	List() []string
}

func List(lister Lister, logger Logger) []string {
	logger.Debug("Listing all equations...")

	return lister.List()
}
