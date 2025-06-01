package main

import (
	"database/sql"
	"log"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/internal/routers"
	 _ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := repo.NewStore(conn)
	server := routers.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}