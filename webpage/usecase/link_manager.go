package usecase

type LinkManager struct {
}

func (l *LinkManager) CategorizeLinks(url string, links []string) ([]string, []string, error) {

	internalLinks := make([]string, 0, 100)
	externalLinks := make([]string, 0, 100)

	return internalLinks, externalLinks, nil
}
