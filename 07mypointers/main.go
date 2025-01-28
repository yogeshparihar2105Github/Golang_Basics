package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("Value of not initialized ptr : ", ptr) // default value is nil

	var mynumber int = 4

	ptr = &mynumber

	fmt.Println("mynumber value before update :", mynumber)
	*ptr = *ptr + 5 // update the value in myNumber
	fmt.Println("mynumber value after update :", mynumber)

	fmt.Println("Value of Pointer is :", ptr)
	fmt.Println("Value where Pointer is pointing :", *ptr)
}
