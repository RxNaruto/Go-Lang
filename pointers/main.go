package main

import "fmt"

func main() {
	myNumber := 23
	//var ptr *int
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ", *ptr)
}
