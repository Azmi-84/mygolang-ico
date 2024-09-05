package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Welcome message
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	// Create a new reader
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the user's name
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Thanks for the input,", name)

	// Prompt for the user's age
	fmt.Print("Enter your age: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}
	fmt.Println("Thanks for the input,", age)
	fmt.Printf("Type of input is %T\n", age)

	// Prompt for a number to check if it is even or odd
	var n int
	fmt.Print("Enter a number: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Check if the number is even or odd
	if ok, err := isEven(n); ok {
		fmt.Println("it's an even number.")
	} else {
		fmt.Println(err)
	}
}

// isEven checks if a number is even or odd.
func isEven(n int) (bool, error) {
	if n&1 == 1 {
		return false,
			fmt.Errorf("it's an odd number")
	}
	return true, nil
}
