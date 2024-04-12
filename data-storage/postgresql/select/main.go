package main

import (
	"database/sql"
	"fmt"
	"practice_go/data-storage/postgresql/model"
)

func teacher(id int, db *sql.DB) (*model.Teacher, error) {
	teacher := model.Teacher{}
	err := db.QueryRow("SELECT id, firstname, lastname FROM teacher WHERE id > $1", id).
		Scan(&teacher.ID, &teacher.Firstname, &teacher.Lastname)
	if err != nil {
		fmt.Printf("failed to select: %v", err)
		return nil, err
	}
	return &teacher, nil
}
