package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	result := doSum()
	fmt.Println(result)

	f, err := os.Create("profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
}

// go tool pprof gettingstarted profile.pb.gz
// gunzip profile.pb.gz -- удаляет profile.pb.gz и создает новый файл profile.pb

func doSum() int {
	sum := 0

	for i := 0; i < 787766777; i++ {
		sum += i
	}

	return sum
}
