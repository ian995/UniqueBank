package tests

import (
	"log"
	"os"
	"testing"
	"database/sql"
	 _ "github.com/lib/pq"

	"github.com/ian995/UniqueBank/internal/repo"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable"
)

var testQueries *repo.Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = repo.New(testDb)

	os.Exit(m.Run())
}