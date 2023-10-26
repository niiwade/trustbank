package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://trustbank:trustapi@localhost:5432/trustbankdb?sslmode=disable"
)

var TestQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the db:", err)
	}

	TestQueries = New(conn)

	os.Exit(m.Run())
}
