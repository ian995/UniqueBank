package tests

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/ian995/UniqueBank/db/sqlc"
	"github.com/jackc/pgx/v5"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	ctx := context.Background()

	
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close(ctx)

	testQueries = db.New(conn)

	os.Exit(m.Run())
}