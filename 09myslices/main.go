package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to my slices")

	var fruitList = []string{"apple", "tomoato", "peach"} // if you don't give the size it will become slice otherwise array

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// --------------------------------------------------------------------------
	highscores := make([]int, 4) // this is slice

	highscores[0] = 234
	highscores[1] = 233
	highscores[2] = 223 //default value is zero
	highscores[3] = 253
	// highscores[4] = 233 // not allowed but you can use append

	highscores = append(highscores, 55, 434, 645)

	fmt.Println(highscores)
	fmt.Printf("Type of highscores is %T \n", highscores)

	fmt.Println("Is sorted :", sort.IntsAreSorted(highscores))

	sort.Ints(highscores)

	fmt.Println("after sorted", highscores)

	//VERY IMPORTANT

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "go", "java", "c", "python", "c#"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
