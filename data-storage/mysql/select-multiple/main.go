package main

import (
	"database/sql"
	"time"
)

type Teacher struct {
	ID          int
	CreatedTime time.Time
	UpdatedTime time.Time
	Firstname   string
	Lastname    string
}

func selectMultiple(db *sql.DB) (*[]Teacher, error) {
	rows, err := db.Query("SELECT id, firstname, lastname FROM teacher")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teachers := make([]Teacher, 0)

	for rows.Next() {
		teacher := Teacher{}
		if err := rows.Scan(&teacher.ID, &teacher.Firstname, &teacher.Lastname); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return &teachers, nil
}
