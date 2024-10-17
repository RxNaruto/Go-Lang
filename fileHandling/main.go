package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is the text for file"
	file, err := os.Create("./myfiles.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is :", length)
	defer file.Close()
	readFile("myfiles.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text inside the file is :\n", databyte)
	fmt.Println("Text inside the file is :\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
