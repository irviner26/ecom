package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/irviner26/ecom/config"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open(
		"pgx",
		"postgres://"+
			config.Global.DBUser+":"+config.Global.DBPassword+
			"@"+config.Global.DBAddress+"/"+
			config.Global.DBName,
	)
	if err != nil {
		log.Fatal(err)

		return
	}
	defer db.Close()

	driver, err := pgx.WithInstance(db, &pgx.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"pgx",
		driver,
	)
	if err != nil {
		log.Fatal(err)

		return
	}

	cmd := os.Args[(len(os.Args))-1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)

			return
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)

			return
		}
	}
}
