package environment

import "github.com/jmoiron/sqlx"

type repositoryManager interface {
	GetByName(name string) (*Environment, error)
	Create(environment Environment) error
}

type repository struct {
	db *sqlx.DB
}

func (r repository) Create(environment Environment) error {
	tx := r.db.MustBegin()
	tx.NamedExec("INSERT INTO environment (name) VALUES (:name)", environment)
	return tx.Commit()
}

func (r repository) GetByName(name string) (*Environment, error) {
	env := Environment{}
	err := r.db.Get(&env, "SELECT * FROM environment WHERE name = $1", name)
	return &env, err
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}
