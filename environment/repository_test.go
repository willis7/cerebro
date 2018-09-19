package environment

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willis7/cerebro/test"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func setupRepoTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repository) {
	mockDB, mock, sqlxDB := test.NewMockSqlxDB(t)
	repo := NewRepository(sqlxDB)
	return mockDB, mock, *repo
}

func TestRepository_Create(t *testing.T) {
	mockDB, mock, repo := setupRepoTest(t)
	defer mockDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO environment").WithArgs("test store")
	mock.ExpectCommit()

	repo.Create(Environment{Name: "test store"})

	test.AssertExpectationsMet(mock, t)
}

func TestNewRepository_GetByID(t *testing.T) {
	mockDB, mock, repo := setupRepoTest(t)
	defer mockDB.Close()

	columns := []string{"id", "name"}
	name := "test store"
	mock.ExpectQuery("SELECT *").
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, name))

	env, _ := repo.GetByID(1)

	expect := &Environment{
		ID:   1,
		Name: name,
	}

	assert.Equal(t, expect, env)
	test.AssertExpectationsMet(mock, t)
}

func TestRepository_GetByName(t *testing.T) {
	mockDB, mock, repo := setupRepoTest(t)
	defer mockDB.Close()

	columns := []string{"id", "name"}
	name := "test store"
	mock.ExpectQuery("SELECT *").
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(1, name))

	env, _ := repo.GetByName(name)

	expect := &Environment{
		ID:   1,
		Name: name,
	}

	assert.Equal(t, expect, env)
	test.AssertExpectationsMet(mock, t)
}

//func TestRepository_GetByName_Invalid(t *testing.T) {
//	store := mockStore{}
//	repo := repository{
//		db: store,
//	}
//
//	_, actual := repo.GetByName("invalid")
//	expected := errors.New("environment not found")
//
//	if actual.Error() != expected.Error() {
//		t.Fail()
//	}
//}
