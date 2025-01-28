package main

import "fmt"

func main() {
	fmt.Println("Welcome to function")
	greeter()
	greeterTwo()

	// result := addr(2, 4)

	result := proAdder(3, 5, 6, 5)

	fmt.Println(result)
}

func addr(valone, valtwo int) (int, string) { //when returning two values use ()
	return valone + valtwo, "hii"
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Namastey from Golang")
}

func greeterTwo() {
	fmt.Println("Greeter Two")
}
