package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbUri    = "postgresql://root:TEst.0429.30@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries

// load env variables
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbdriver, dbUri)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
