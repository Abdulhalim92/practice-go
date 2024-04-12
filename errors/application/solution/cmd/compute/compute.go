package main

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"log"
	"os"
	"practice_go/errors/application/solution/invoice"
	"time"
)

const (
	startDate   = "2003-07-01"
	endDate     = "2020-07-01"
	baseDirPath = "./errors/application/solution/data"
)

var (
	totalVat    *money.Money
	totalVatExc *money.Money
)

func init() {
	totalVat = money.New(0, money.USD)
	totalVatExc = money.New(0, money.USD)
}

func main() {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Fatalf("impossible to parse start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Fatalf("impossible to parse end date: %v", err)
	}

	// from start to end date add 1 month at each iteration
	for d := start; d.Unix() < end.Unix(); d = d.AddDate(0, 1, 0) {
		// create a var that will contain the name of the dir to open
		monthDirPath := fmt.Sprintf("%s/%s-%d", baseDirPath, d.Month(), d.Year())
		// open the directory
		dir, err := os.Open(monthDirPath)
		if err != nil {
			log.Fatalf("failed to open directory: %s: %s", monthDirPath, err)
		}
		// defer the closing of the directory
		defer dir.Close()
		//read all files and folder into the dir
		list, err := dir.Readdirnames(0)
		if err != nil {
			log.Fatalf("failed to read all files in directory: %s: %s", monthDirPath, err)
		}
		// iterate for each filename in list
		for _, name := range list {
			// construct dir path
			filePath := fmt.Sprintf("%s/%s", monthDirPath, name)
			// extract data from dir
			vatExc, vat, err := invoice.ReadFromFile(filePath)
			if err != nil {
				log.Fatalf("failed to parse invoice %s: %s", filePath, err)
			}
			totalVat, err = totalVat.Add(vat)
			if err != nil {
				log.Fatalf("impossible to add VAT to counter")
			}
			totalVatExc, err = totalVatExc.Add(vatExc)
			if err != nil {
				log.Fatalf("impossible to add VAT to counter")
			}
		}
	}

	log.Println("Total VAT: ", totalVat.Display())
	log.Println("Total VAT Exc: ", totalVatExc.Display())
}
