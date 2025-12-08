package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresStorage(dbAddr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Opened connection to database (PostgreSQL)")

	return db, err
}