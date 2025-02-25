package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Wexlersolk/Grief/internal/store"
)

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: fmt.Sprintf("user%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Role: store.Role{
				Name: "user",
			},
		}
	}

	return users
}
