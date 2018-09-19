package environment

type Service struct {
	repo repositoryManager
}

func NewService(repo repositoryManager) *Service {
	return &Service{
		repo: repo,
	}
}

func (s Service) GetByName(name string) (*Environment, error) {
	return s.repo.GetByName(name)
}

func (s Service) Create(env Environment) error {
	return s.repo.Create(env)
}
