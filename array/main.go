package main

import "fmt"

func main() {
	fmt.Println("Array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))
	//[Apple Banana  Mango] - the space here indicates that there is a value that is missing
	var veglist = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("veg list is: ", veglist)

}
