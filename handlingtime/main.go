package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is time study for golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// if you want to get the present time in specific format
	fmt.Println(presentTime.Format("01-02-2006 "))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04"))

	//if you want to create the date
	createdDate := time.Date(2023, time.August, 5, 15, 30, 00, 00, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04"))
}
