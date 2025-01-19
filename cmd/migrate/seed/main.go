package main

import (
	"log"

	"github.com/Wexlersolk/Grief/internal/db"
	"github.com/Wexlersolk/Grief/internal/env"
	"github.com/Wexlersolk/Grief/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/db?sslmode=disable")
	log.Println("DB Connection Address:", addr)
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("new conn created")

	defer conn.Close()

	store := store.NewStorage(conn)
	log.Println("new store created")

	db.Seed(store, conn)

	log.Println("new seed created")
}
