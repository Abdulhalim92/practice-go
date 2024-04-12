package main

import (
	"encoding/xml"
	"fmt"
)

type Product struct {
	Name string `xml:"first>second>third"`
}

func main() {
	c := Product{Name: "testing"}

	b, err := xml.MarshalIndent(c, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

/*
<Product>
        <first>
                <second>
                        <third>testing</third>
                </second>
        </first>
</Product>
*/
