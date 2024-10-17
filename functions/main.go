package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Hi there from main")
	//you cannot create function inside function
	result := adder(1, 2)
	fmt.Println(result)
	proRes, proMessage := proAdder(236, 6, 7, 4, 6)
	fmt.Println(proRes)
	fmt.Println(proMessage)

}

func greeter() {
	fmt.Println("Hi there from greeter")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

// pro functions
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "this is message"
}
