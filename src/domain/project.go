package domain

type Project struct {
	Id       int
	Name     string
	Provider Provider
	VcsUrl   string
}

func (p Project) GetProvider() Provider {
	return p.Provider
}
