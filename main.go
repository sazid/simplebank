package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sazid/simplebank/api"
	db "github.com/sazid/simplebank/db/sqlc"
)

const (
	// dbDriver = "postgres"
	dbSource      = "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8300"
)

func main() {
	var err error
	conn, err := pgxpool.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
