package main

import "fmt"

const loginToken string = "auth" //public

func main() {
	var username string = "naruto"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint = 16
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 16.234234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var newVal int
	fmt.Println(newVal)
	fmt.Printf("Variable is of type: %T \n", newVal)

	numberOfUser := 30000 //no var style
	fmt.Println(numberOfUser)

}
