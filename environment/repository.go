package environment

type storer interface {
	Persist(environment Environment) error
	RetrieveByName(name string) (*Environment, error)
}

type repositoryManager interface {
	GetByName(name string) (*Environment, error)
	Create(env Environment) error
}

type repository struct {
	store storer
}

func (r repository) GetByName(name string) (*Environment, error) {
	return r.store.RetrieveByName(name)
}

func (r repository) Create(env Environment) error {
	return nil
}
