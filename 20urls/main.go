package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/search?q=yogesh&email=yogeshparihar2105@gmail.com"

func main() {
	fmt.Println("Welcome to handling urls in Golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println("Result : ", result)
	// fmt.Println("Scheme : ", result.Scheme)
	// fmt.Println("Host : ", result.Host)
	// fmt.Println("User : ", result.User)
	// fmt.Println("Path : ", result.Path)
	// fmt.Println("Raw Query : ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The Type of query parama are : %T\n", qparams)

	// fmt.Println(qparams["q"])
	// fmt.Println(qparams["email"])

	for _, val := range qparams {
		fmt.Println("param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=yogesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
