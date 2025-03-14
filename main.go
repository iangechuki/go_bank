package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/iangechuki/go_bank/api"
	db "github.com/iangechuki/go_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://admin:adminpassword@localhost:5432/go_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
