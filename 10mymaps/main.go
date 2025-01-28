package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["JV"] = "Java"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages :", languages)

	fmt.Println("JS is short name for", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages after delete :", languages)

	//looping in maps
	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}

}
