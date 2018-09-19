package environment

import (
	"testing"

	"github.com/willis7/cerebro/test"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRepository_Create(t *testing.T) {
	mockDB, mock, sqlxDB := test.NewMockSqlxDB(t)
	defer mockDB.Close()

	repo := repository{
		db: sqlxDB,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO environment").WithArgs("test store")
	mock.ExpectCommit()

	repo.Create(Environment{Name: "test store"})

	assertExpectationsMet(mock, t)
}

func TestRepository_GetByName(t *testing.T) {
	mockDB, mock, sqlxDB := test.NewMockSqlxDB(t)
	defer mockDB.Close()

	repo := repository{
		db: sqlxDB,
	}

	columns := []string{"name"}
	expected := "test store"

	mock.ExpectQuery("SELECT *").
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow(expected))

	env, _ := repo.GetByName(expected)
	if env.Name != expected {
		t.Fatalf("expected: %s, got: %s", expected, env.Name)
	}

	assertExpectationsMet(mock, t)
}

func assertExpectationsMet(mock sqlmock.Sqlmock, t *testing.T) {
	// ensure all expectations have been met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error: %s", err)
	}
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
