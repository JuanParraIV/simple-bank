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
	dbUri    = "postgresql://root:TEst.0429.30@localhost:5432/simple_bank?sslmode=disable"
)
var testDB *sql.DB
var testQueries *Queries

// load env variables
// TestMain is the entry point for testing in this package. It sets up the 
// database connection required for the tests and initializes the testQueries 
// object used to interact with the database. If the database connection 
// cannot be established, the function logs the error and terminates the 
// program. After running all tests, it exits with the appropriate status code.
func TestMain(m *testing.M) {
	// for testing purposes, we are using a hardcoded database
	// with a specific connection string
	// Load environment variables
	//config, err := utils.LoadConfig("../..")

	//if err != nil {
	//	log.Fatal("cannot load config:", err)
	//}
	//testDB, err = sql.Open(config.DBDriver, config.DBSource)
	var err error
	
	testDB, err = sql.Open(dbdriver, dbUri)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
