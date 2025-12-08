package main

import (
	"database/sql"
	"log"

	"github.com/irviner26/ecom/cmd/api"
	"github.com/irviner26/ecom/config"
	"github.com/irviner26/ecom/db"
)

func main() {
	db, err := db.NewPostgresStorage(
		"postgres://" +
		config.Global.DBUser + ":" + config.Global.DBPassword +
		"@" + config.Global.DBAddress + "/" + 
		config.Global.DBName,
	)

	initStorage(db)

	server := api.NewAPIServer(config.Global.PublicHost + ":" + config.Global.Port, db)
	if err = server.Run(); err != nil {
		log.Fatalln(err)
	}
}

func initStorage (db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to database")
}