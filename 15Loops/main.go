package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Loop Learning!")

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println("Original order of days:", days)

	for i, k := 0, len(days)-1; i < k; i, k = i+1, k-1 {
		days[i], days[k] = days[k], days[i]
	}
	fmt.Println("Reversed order of days:", days)

	sort.Strings(days)
	fmt.Println("Alphabetical order of days:", days)

	for m, day := range days {
		fmt.Printf("%v := %v\n", m, day)
	}

	// Multiplication Table in Reverse order

	for i := 1; i < 11; i++ {
		fmt.Printf("Multiplication table of %d:\n", i)
		for k := 1; k < 11; k++ {
			var multiplicationTable int = i * k
			fmt.Printf(" %d * %d = %d\n", i, k, multiplicationTable)
		}
		fmt.Println()
	}

}
