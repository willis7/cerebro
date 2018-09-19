package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/willis7/cerebro/environment"
)

func main() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=postgres dbname=cerebro sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	repo := environment.NewRepository(db)
	_ = environment.NewService(repo)
}
