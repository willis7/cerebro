package test

import (
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func NewMockSqlxDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *sqlx.DB) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Should not get an error in creating a mock database")
	}
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	sqlxDB.SetMaxOpenConns(10)
	return mockDB, mock, sqlxDB
}
