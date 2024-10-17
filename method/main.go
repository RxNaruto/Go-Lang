package main

import "fmt"

func main() {
	//no inheritance in golang
	naruto := User{"naruto", "naruto@hokagemail.com", true, 19}
	fmt.Println(naruto)
	fmt.Printf("Naruto details are: %+v\n", naruto)
	//if you want to get specific value
	fmt.Printf("Name is %v and Email is %v\n", naruto.Name, naruto.Email)
	naruto.GetStatus()
	naruto.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// method creating
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New mail is: ", u.Email)
}
