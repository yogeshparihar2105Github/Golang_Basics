package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang array")

	var fruitList [5]string // you have to mention the size of array

	fruitList[0] = "Apple"
	fruitList[1] = "Tomota"
	// it keeps blank space for the empty array element
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is :", fruitList)
	fmt.Println("Length Fruit List is :", len(fruitList))

	var veglist = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Veggy list is :", veglist)
}
