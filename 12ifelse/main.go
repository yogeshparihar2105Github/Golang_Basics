package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Not Regular user"
	} else {
		result = "Suspect because it is 10"
	}

	fmt.Println(result)

}
