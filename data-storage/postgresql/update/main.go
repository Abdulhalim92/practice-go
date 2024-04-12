package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func updateTeacher(id int, db *sql.DB) error {
	// обновить имя учителя
	res, err := db.Exec("UPDATE teacher SET firstname = 'Hendrick' WHERE id = $1", id)
	if err != nil {
		fmt.Printf("failed to update: %v", err)
		return err
	}

	// получить количество обновленных строк
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("failed to get rows affected: %v", err)
		return err
	}

	if affected != 1 {
		fmt.Printf("Something went wrong %d rows were affected\n", affected)
		return errors.New("something went wrong")
	}

	fmt.Printf("Update is a success")

	return nil
}
