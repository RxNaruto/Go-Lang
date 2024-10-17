package main

import "fmt"

func main() {
	//make(map[key]value)
	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["ts"] = "typescript"

	fmt.Println("List of languages: ", languages)
	fmt.Println("Js shorts for: ", languages["js"])

	//to delete a value

	delete(languages, "py")
	fmt.Println("The languages are: ", languages)
	for key, value := range languages {
		fmt.Println(key, value)
	}

}
