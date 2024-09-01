package main

import "fmt"

func main() {
	// var Array [5]int

	// Array[0] = 10
	// Array[1] = 20
	// Array[2] = 30
	// Array[3] = 40

	// fmt.Println("Value of Array is: ", Array)
	// fmt.Println("Length of Array is: ", len(Array)) // Length of Array is:  5. this is because we have defined the length of the array as 5. no matter how many elements we have added to the array

	// Another way to define an array
	var Array_00 = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Value of Array is: ", Array_00)
}
