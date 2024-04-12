package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = create(db)
	if err != nil {
		fmt.Printf("failed to create: %v", err)
		return
	}
}

func create(db *sql.DB) error {
	file, err := os.Open("data-storage/postgresql/schemas/init.sql")
	if err != nil {
		log.Printf("failed to open file: %v", err)
		return err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("failed to read file: %v", err)
		return err
	}

	res, err := db.Exec(string(b))
	if err != nil {
		fmt.Printf("failed to execute query: %v", err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("failed to get last insert id: %v", err)
		return err
	}

	fmt.Printf("last insert id: %d\n", id)

	return nil
}
