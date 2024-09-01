package main

import "fmt"

func main() {
	// var pointer *int
	// fmt.Println("Value of pointer is: ", pointer)

	number := 10
	pointer_00 := &number
	fmt.Println("Value of number is: ", pointer_00)
	fmt.Println("Value of number is: ", *pointer_00)

	*pointer_00 = *pointer_00 * 10
	fmt.Println("Value of number is: ", number)
	// fmt.Println("Value of number is: ", *pointer_00)
}
