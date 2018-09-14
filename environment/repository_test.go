package environment

import (
	"errors"
	"testing"
)

type mockStore struct{}

func (s mockStore) Persist(environment Environment) error {
	return nil
}

func (s mockStore) RetrieveByName(name string) (*Environment, error) {
	if name == "invalid" {
		return &Environment{}, errors.New("environment not found")
	} else {
		return &Environment{Name: "test store"}, nil
	}
}

func TestRepository_CreateCreate(t *testing.T) {
	//	given a repository with a
	m := mockStore{}
	repo := repository{
		store: m,
	}

	//	when I call the Create method
	actual := repo.Create(Environment{})

	if actual != nil {
		t.Fail()
	}
}

func TestRepository_GetByName(t *testing.T) {
	store := mockStore{}
	repo := repository{
		store: store,
	}

	actual, _ := repo.GetByName("test store")
	expected := "test store"

	if actual.Name != expected {
		t.Fail()
	}
}

func TestRepository_GetByName_Invalid(t *testing.T) {
	store := mockStore{}
	repo := repository{
		store: store,
	}

	_, actual := repo.GetByName("invalid")
	expected := errors.New("environment not found")

	if actual.Error() != expected.Error() {
		t.Fail()
	}
}
