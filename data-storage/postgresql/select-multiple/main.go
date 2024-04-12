package main

import (
	"database/sql"
	"practice_go/data-storage/postgresql/model"
)

func selectMultiple(db *sql.DB) (*[]model.Teacher, error) {
	rows, err := db.Query("SELECT id, firstname, lastname FROM teacher")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teachers := make([]model.Teacher, 0)

	for rows.Next() {
		teacher := model.Teacher{}
		if err := rows.Scan(&teacher.ID, &teacher.Firstname, &teacher.Lastname); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return &teachers, nil
}
