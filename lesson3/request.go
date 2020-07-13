package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), nil

}

func main() {

	type1, err := getContentType("https://www.google.com/")
	if err != nil {
		fmt.Println("Type 1 is error ")
		fmt.Println(err)
	} else {
		fmt.Println("Type 1 response")
		fmt.Println(type1)
	}

	type2, err := getContentType("https://foo.baz.com/")
	if err != nil {
		fmt.Println("Type 2 is error ")
		fmt.Println(err)
	} else {
		fmt.Println("Type 2 response")
		fmt.Println(type2)
	}

}
