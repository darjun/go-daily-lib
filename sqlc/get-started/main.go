package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/net/context"

	"github.com/darjun/go-daily-lib/sqlc/get-started/db"
)

func main() {
	pq, err := sql.Open("postgres", "dbname=sqlc sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(pq)

	authors, err := queries.ListAuthors(context.Background())
	if err != nil {
		log.Fatal("ListAuthors error:", err)
	}
	fmt.Println(authors)

	insertedAuthor, err := queries.CreateAuthor(context.Background(), db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		log.Fatal("CreateAuthor error:", err)
	}
	fmt.Println(insertedAuthor)

	fetchedAuthor, err := queries.GetAuthor(context.Background(), insertedAuthor.ID)
	if err != nil {
		log.Fatal("GetAuthor error:", err)
	}
	fmt.Println(fetchedAuthor)

	err = queries.DeleteAuthor(context.Background(), insertedAuthor.ID)
	if err != nil {
		log.Fatal("DeleteAuthor error:", err)
	}
}
