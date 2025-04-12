package main

import (
	"database/sql"
	"log"

	"github.com/juanparraiv/simple-bank/api"
	db "github.com/juanparraiv/simple-bank/db/sqlc"
	"github.com/juanparraiv/simple-bank/utils"
	_ "github.com/lib/pq"

)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// Create a new server instance
	store := db.NewStore(conn)
	server := api.NewServer(store)

	// Start the server on port 9001
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
