package adapter

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/choonhong/hotel-data-merge/ent"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*ent.Client, error) {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	fmt.Println("Connected to sqlite")

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
