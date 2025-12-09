package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgresStorage(dbAddr string, ctx context.Context) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, dbAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Opened connection to database (PostgreSQL)")

	return db, err
}
