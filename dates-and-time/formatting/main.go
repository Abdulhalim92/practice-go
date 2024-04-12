package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// Форматирование времени
	fmt.Println(now.Format("Mon Jan 2 "))
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println(now.Format(time.RFC3339))
}

// Mon Jan 2 15:04:05 -0700 MST 2006   ---   базовая дата
