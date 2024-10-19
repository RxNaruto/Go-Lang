package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummyjson.com/c/3029-d29f-4014-9fb4"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of response is %T\n", response)

	defer response.Body.Close() //caller reponsibility

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}
