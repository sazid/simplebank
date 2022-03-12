package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sazid/simplebank/util"
)

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err = pgxpool.Connect(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
