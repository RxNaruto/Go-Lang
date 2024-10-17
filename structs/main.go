package main

import "fmt"

func main() {
	//no inheritance in golang
	naruto := User{"naruto", "naruto@hokagemail.com", true, 19}
	fmt.Println(naruto)
	fmt.Printf("Naruto details are: %+v\n", naruto)
	//if you want to get specific value
	fmt.Printf("Name is %v and Email is %v", naruto.Name, naruto.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
