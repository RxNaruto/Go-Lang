package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	PerformGetRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length is :", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	//other way of doing this

}
