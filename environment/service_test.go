package environment

import (
	"errors"
	"testing"
)

type RepositoryMock struct{}

func (r RepositoryMock) GetByName(name string) (*Environment, error) {
	if name == "test" {
		return &Environment{Name: "test"}, nil
	} else {
		return &Environment{}, errors.New("name not found")
	}

}

func (r RepositoryMock) GetByID(id int) (*Environment, error) {
	if id == 1 {
		return &Environment{ID: 1}, nil
	} else {
		return &Environment{}, errors.New("id not found")
	}

}

func (r RepositoryMock) Create(env Environment) error {
	return nil
}

func TestService_GetEnvironmentByName(t *testing.T) {
	mockR := RepositoryMock{}
	svc := NewService(mockR)

	actual, _ := svc.GetByName("test")
	expected := "test"

	//  expect the correct value is returned for name
	if actual.Name != expected {
		t.Fail()
	}
}

func TestSevice_Create(t *testing.T) {
	mockR := RepositoryMock{}
	svc := NewService(mockR)
	env := Environment{Name: "test create"}

	actual := svc.Create(env)

	if actual != nil {
		t.Fail()
	}
}
