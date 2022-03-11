package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

const (
	// dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
