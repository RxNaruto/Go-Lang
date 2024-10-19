package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://rithktech.dev:3000/learn?coursename=reactjs&paymentid=abt3b523g"

func main() {

	fmt.Println("Url operations")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of querry params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	//constructing a url

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/newgo",
		RawPath: "user=naruto",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
