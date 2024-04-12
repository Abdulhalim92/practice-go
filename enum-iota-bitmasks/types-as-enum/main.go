package main

import (
	"encoding/json"
	"fmt"
)

type HTTPMethod int

const (
	GET     HTTPMethod = 0
	POST    HTTPMethod = 1
	PUT     HTTPMethod = 2
	DELETE  HTTPMethod = 3
	PATCH   HTTPMethod = 4
	HEAD    HTTPMethod = 5
	OPTIONS HTTPMethod = 6
	TRACE   HTTPMethod = 7
	CONNECT HTTPMethod = 8
)

type HTTPRequest struct {
	Method  HTTPMethod
	Headers map[string]string
	Uri     string
}

func handle(method HTTPMethod, headers map[string]string, uri string) {
	if method == GET {
		fmt.Println("the Method is get")
	} else {
		fmt.Println("the Method is not get")
	}
}

func main() {
	r := HTTPRequest{
		Method:  GET,
		Headers: map[string]string{"Accept": "application/json"},
		Uri:     "/prices",
	}
	fmt.Println(r)

	data, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	jsonB := []byte("{\"Method\":\"GET\",\"Headers\":{\"Accept\":\"application/json\"},\"Uri\":\"/prices\"}")
	req := HTTPRequest{}
	err = json.Unmarshal(jsonB, &req)
	if err != nil {
		panic(err)
	}
}
