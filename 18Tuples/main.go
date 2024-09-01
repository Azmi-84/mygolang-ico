package main

import (
	"fmt"
)

func exponents(x int) (int, int) {
	return x * x, x * x * x
}

func main() {
	var square int
	var cube int
	var number int = 15
	square, cube = exponents(number)
	fmt.Printf("Square of %v is: %v\n", number, square)
	fmt.Printf("Cube of %v is: %v\n", number, cube)
}
