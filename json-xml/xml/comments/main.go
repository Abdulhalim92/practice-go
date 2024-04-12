package main

import (
	"encoding/xml"
	"fmt"
)

type Price struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Product struct {
	Comment string `xml:",comment"`
	Price   Price  `xml:"price"`
	Name    string `xml:"name"`
}

func main() {
	c := Product{
		Comment: "this is my comment",
		Price: Price{
			Text:     "123",
			Currency: "EUR",
		},
		Name: "testing",
	}

	b, err := xml.MarshalIndent(c, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

/*
<Product>
        <!--this is my comment-->
        <price currency="EUR">123</price>
        <name>testing</name>
</Product>
*/
