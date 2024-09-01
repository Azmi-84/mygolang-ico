package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Struct Learning")

	Azmi := User{"azmi-84", 12, "azmi@go.com", true}
	fmt.Println(Azmi)
	fmt.Printf("Details about Azmi: %+v\n", Azmi)
	fmt.Printf("Name is %v and email is %v\n.", Azmi.Name, Azmi.Email)
}

type User struct {
	Name       string
	Age        int
	Email      string
	isVerified bool
}
