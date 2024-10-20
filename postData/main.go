package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformPostReq()

}

func PerformPostReq() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
	 {
	    "coursename": "Let's go with golang",
		"price": 0 
	 }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
