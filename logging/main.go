package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file")
	}
	defer func() { _ = f.Close() }()

	mw := io.MultiWriter(f, os.Stdout)
	log.SetOutput(mw)

	log.Println("System is starting")
	log.Println("Retrieving user 12")
	log.Println("Compute the total of invoice 1262663663")

}

//

// panic()
// Когда регистратор звонит panic от вашего имени:
// Любые отложенные функции, определенные в текущей функции, будут вызываться
// Текущая функция вернется к вызывающей стороне
// В вызывающей стороне будут вызваны все определенные отложенные функции, и
// вызывающая сторона вернется к функции, которая ее вызвала...
// Вызов panic вызывает все отложенные функции, определенные в стеке
// вызовов горутины.
