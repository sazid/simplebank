package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sazid/simplebank/api"
	db "github.com/sazid/simplebank/db/sqlc"
	"github.com/sazid/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := pgxpool.Connect(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
