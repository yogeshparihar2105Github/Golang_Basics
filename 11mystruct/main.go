package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	//no inheritance in golang; No super or parent
	//don't have to look in multiple files
	//you should be able to understand whatever on screen

	yogesh := User{
		"Yogesh",
		"yogeshparihar2105@gmail.com",
		true,
		26,
	}

	fmt.Println(yogesh)
	fmt.Printf("yogesh details are : %+v\n", yogesh) //in struct use %+v for entire details
	fmt.Printf("Name is %v \nEmail is %v", yogesh.Name, yogesh.Email)
}

//Capital to make it public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
