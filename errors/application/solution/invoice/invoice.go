package invoice

import (
	"encoding/csv"
	"errors"
	"github.com/Rhymond/go-money"
	"io"
	"os"
	"strconv"
)

func ReadFromFile(filename string) (*money.Money, *money.Money, error) {
	record, err := readCSV(filename)
	if err != nil {
		return nil, nil, err
	}
	// record: invoice number [0326582789 UnicornQuiver 126730 25346 152076 5]

	vatExcConverted, err := strconv.Atoi(record[2])
	if err != nil {
		return nil, nil, err
	}
	vatExc := money.New(int64(vatExcConverted), money.USD)

	varConverted, err := strconv.Atoi(record[3])
	if err != nil {
		return nil, nil, err
	}
	vat := money.New(int64(varConverted), money.USD)

	return vatExc, vat, nil
}

func readCSV(filename string) ([]string, error) {
	invoiceCVS, err := os.Open(filename)
	defer func(invoiceCVS *os.File) { _ = invoiceCVS.Close() }(invoiceCVS)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(invoiceCVS)
	record, err := reader.Read()
	if err == io.EOF {
		return nil, errors.New("file is empty")
	}
	if err != nil {
		return nil, err
	}

	return record, nil
}
