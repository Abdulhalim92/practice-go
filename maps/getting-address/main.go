package main

type emlpoyee struct {
	name     string
	genre    string
	position string
}

func main() {
	employees := make(map[string]emlpoyee)

	employees["first"] = emlpoyee{
		name:     "John Smith",
		genre:    "M",
		position: "CEO",
	}

	//fmt.Printf("address ot the %p", &employees["first"])
	// Cannot take the address of 'employees["first"]'
}
