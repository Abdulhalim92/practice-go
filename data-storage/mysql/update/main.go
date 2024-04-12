package main

import (
	"database/sql"
	"fmt"
)

func update(firstname, lastname string, id int, db *sql.DB) error {
	res, err := db.Exec("UPDATE teacher SET firstname = ?, lastname = ? WHERE id = ?", firstname, lastname, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// получить количество обновленных строк
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if affected != 1 {
		fmt.Printf("Something went wrong %d rows were affected\n", affected)
	} else {
		fmt.Println("Update is a success")
	}

	return nil
}
