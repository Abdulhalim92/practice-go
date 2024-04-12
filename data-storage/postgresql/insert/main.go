package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func createTeacher(firstname string, lastname string, db *sql.DB) (int, error) {
	insertedID := 0
	err := db.QueryRow(
		"INSERT INTO public.teacher (`create_time`, `firstname`, `lastname`)"+
			" VALUES (NOW(), $1, $2) RETURNING id", firstname, lastname).
		Scan(&insertedID)
	if err != nil {
		fmt.Printf("failed to insert: %v", err)
		return 0, err
	}
	if insertedID == 0 {
		return 0, errors.New("something went wrong id inserted is equal to zero")
	}

	return insertedID, nil
}
