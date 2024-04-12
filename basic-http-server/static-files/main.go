package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Задаем директорию, из которой будут отдаваться статические файлы
	staticDir := "./basic-http-server/static-files/static"

	// Проверяем, существует ли директория с файлами
	_, err := os.Stat(staticDir)
	if os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", staticDir)
		return
	}

	// Настройка обработчика для статических файлов
	fs := http.FileServer(http.Dir(staticDir))

	// Устанавливаем маршрут для обработки статических файлов
	http.Handle("/", fs)

	// Устанавливаем обработчик ошибок для логгирования ошибок сервер
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Ошибка сервера: Internal Server Error")
	})

	// Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %s\n", err)
	}
}
