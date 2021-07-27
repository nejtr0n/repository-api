package domain

func NewService(providerFactory ProviderFactory) Service {
	return Service{providerFactory: providerFactory}
}

type Service struct {
	providerFactory ProviderFactory
}

func (s Service) GetProject(providerType ProviderType, id int) (*Project, error) {
	provider, err := s.providerFactory.GetProvider(providerType)
	if err != nil {
		return nil, err
	}
	return provider.GetProjectById(id)
}

func (s Service) CreateProject(providerType ProviderType, name string, groupId int) (*Project, error) {
	provider, err := s.providerFactory.GetProvider(providerType)
	if err != nil {
		return nil, err
	}
	return provider.CreateProject(name, groupId)
}

func (s Service) CloneProject(from *Project, to *Project) error {
	fromRepository, err := from.GetProvider().GetVcs().CloneRepository(from)
	if err != nil {
		return err
	}

	err = to.GetProvider().GetVcs().PushRepositoryToProject(to, fromRepository)
	if err != nil {
		return err
	}

	return nil
}
