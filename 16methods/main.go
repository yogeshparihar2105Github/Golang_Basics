package main

import "fmt"

func main() {
	yogesh := User{
		"Yogesh",
		"yogeshparihar2105@gmail.com",
		26,
	}

	yogesh.GetStatus()
	yogesh.NewMail()
	fmt.Println(yogesh.Email)
}

type User struct {
	Name  string
	Email string
	Age   int
}

//user * to pass the actual object and not a copy of it
func (u *User) GetStatus() {
	fmt.Println("User Age :", u.Age)
}

func (u *User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}
