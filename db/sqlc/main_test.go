package db

// Setting up database connection
import (
	"database/sql"
	"testing"

	/**This package provides a generic interface for interacting with SQL databases.
	It still needs to speak with a specific database drive implementation**/
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Define test query object as a global variable
var testQueries *Queries

// This should be comming from ENV files latter

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5439/simple_bank?sslmode=disable"
)

var testDB *sql.DB

// TestMain is the entry point of all unit tests inside one specific Golang package

func TestMain(m *testing.M) {

	var err error

	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot to connect to db", err)
	}

	// Creating a new connection object
	testQueries = New(testDB)

	// Run the unit tests and exits

	os.Exit(m.Run())

}
