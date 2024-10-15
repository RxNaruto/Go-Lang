package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	println("Hello user")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the person:")

	input, _ := reader.ReadString('\n')
	fmt.Println("The name is : ", input)

}
