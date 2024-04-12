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

func createTeacher(firstname string, lastname string, db *sql.DB) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO `teacher` (`create_time`, `firstname`, `lastname`)"+
			" VALUES (NOW(), ?, ?)", firstname, lastname)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func teacher(id int, db *sql.DB) (*Teacher, error) {
	teacher := Teacher{ID: id}
	err := db.QueryRow("SELECT firstname, lastname FROM teacher WHERE id = ?", id).Scan(&teacher.Firstname, &teacher.Lastname)
	if err != nil {
		return &Teacher{}, err
	}
	return &teacher, nil
}
