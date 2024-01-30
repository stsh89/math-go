package equation

type Lister interface {
	List() []string
}

func List(lister Lister, logger Logger) []string {
	logger.Info("Listing all equations...")

	return lister.List()
}
