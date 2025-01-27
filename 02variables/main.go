package main

import "fmt"

// jwt := "324560sdlbos" // You can't use this type of variable declaration syntax in global

const LoginToken string = "sdgyeosdgo" // First Letter Is Capital Means Public

func main() {
	var username string = "Yogesh"
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var smallFloat float64 = 255.4553454563423435
	fmt.Println("Value of float64 : ", smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values as some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type : %T \n", anotherVariable)

	// implicit type

	var website = "google.com"
	fmt.Println(website)
	// website = 3 //not allowed to change the type

	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
}
