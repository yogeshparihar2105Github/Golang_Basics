package main

import "fmt"

func main() {
	defer fmt.Println("World1") //execute at last
	defer fmt.Println("World2") //execute at last before the World1
	defer fmt.Println("World3") //execute at last before the World2
	fmt.Println("Hello")
	defer fmt.Println("World4") //execute at last before the World3
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
