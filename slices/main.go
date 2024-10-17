package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Grapes"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	highScores := make([]int, 4)
	highScores[0] = 200
	highScores[1] = 100
	highScores[2] = 400
	highScores[3] = 300
	highScores = append(highScores, 500, 399, 120)
	fmt.Println(highScores)

	//sorting int
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"react.js", "python", "go", "node.js", "next.js"}
	fmt.Println(courses)

	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
