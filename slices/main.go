package main

import "fmt"

func main() {

	var fruitList = []string{"Apple", "Grapes"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
}
