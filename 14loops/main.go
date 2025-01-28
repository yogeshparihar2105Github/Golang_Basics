package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{
		"Sunday",
		"Tuesday",
		"Wednesday",
		"Friday",
		"Saturday",
	}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, value)
	// }

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco //continue code with lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue //continue skips that particular value (i.e. 5)
			//break 	//break terminate the loops
		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at lco")
}
