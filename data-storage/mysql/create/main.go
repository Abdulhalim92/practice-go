package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/school")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	// открыть файл инициализации таблиц
	f, err := os.Open("./data-storage/mysql/schemas/init.sql")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// читать содержимое файла
	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// выполнить запрос в БД
	res, err := db.Exec(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}
