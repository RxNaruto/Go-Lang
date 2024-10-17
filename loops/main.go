package main

import "fmt"

func main() {
	days := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for key, day := range days {
		fmt.Printf("Index is %v and value is %v\n", key, day)

	}

	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 3 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("new value is: ", rougeValue)
		rougeValue++
	}
lco:
	fmt.Println("Jumping the value")
}
