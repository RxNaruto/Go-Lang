package main

import "fmt"

func main() {

	//defer will run at the last
	//it follows the lifo principle
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
}
