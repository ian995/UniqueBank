package main

import (
	"database/sql"
	"log"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/internal/routers"
	"github.com/ian995/UniqueBank/pkg/utils"
	_ "github.com/lib/pq"
)


func main() {
	config, err := utils.LoadConfig("config/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := repo.NewStore(conn)
	server := routers.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}