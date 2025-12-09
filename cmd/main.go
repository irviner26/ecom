package main

import (
	"context"
	"log"

	"github.com/irviner26/ecom/cmd/api"
	"github.com/irviner26/ecom/config"
	"github.com/irviner26/ecom/db"
	"github.com/jackc/pgx/v5"
)

func main() {
	mainContext := context.Background()

	db, err := db.NewPostgresStorage(
		"postgres://"+
			config.Global.DBUser+":"+config.Global.DBPassword+
			"@"+config.Global.DBAddress+"/"+
			config.Global.DBName,
		mainContext,
	)

	initStorage(db, mainContext)

	server := api.NewAPIServer(config.Global.PublicHost+":"+config.Global.Port, db)
	if err = server.Run(); err != nil {
		log.Fatalln(err)
	}
}

func initStorage(db *pgx.Conn, ctx context.Context) {
	err := db.Ping(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to database")
}
