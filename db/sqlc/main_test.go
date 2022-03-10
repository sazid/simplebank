package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
)

var testQueries *Queries

const (
	// dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
