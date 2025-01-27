package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	for i := 0; i < 15; i++ {
		createdDate := time.Date(2025+i, time.January, 27, 23, 23, 0, 0, time.UTC)
		fmt.Println(createdDate.Format("01-02-2006 Monday"))
	}

}
