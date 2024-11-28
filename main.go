package main

import (
	"context"
	"ginbank/api"
	db "ginbank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	dbSource      = "postgresql://postgres:12345@localhost:5432/microservicegin?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
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
