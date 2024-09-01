package main

import "fmt"

type loginInfo struct {
	username string
	password string
}

func main() {
	fmt.Println("if-else in Go lang")

	login := loginInfo{
		username: "johndoe",
		password: "password123",
	}

	var message string

	if login.username == "" && login.password == "" {
		message = "Please enter both username and password"
	} else if login.username != "johndoe" || login.password != "password123" {
		message = "Invalid username or password"
	} else {
		message = "User successfully logged in"
	}

	fmt.Println(message)
}
