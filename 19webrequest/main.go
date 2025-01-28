package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T", response)
	defer response.Body.Close() //call this as a responsible programmer

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("databyte converted response is ", string(databyte)) //receive the html file of the page
}
