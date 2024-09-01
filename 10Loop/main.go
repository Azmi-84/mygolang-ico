package main

// // import (
// // 	"fmt"
// // 	"math"
// // 	"sort"
// // )

// // func main() {
// // 	// Create a map to store the frequency of each result
// // 	resultFrequency := make(map[int]int)

// // 	// Nested loops to calculate the results and store their frequency
// // 	for index := 0; index <= 10; index++ {
// // 		for index_00 := 0; index_00 < 8; index_00++ {
// // 			for index_01 := 0; index_01 < 6; index_01++ {
// // 				result := int(math.Pow(float64(index), 2) + math.Pow(float64(index_00), 2) + math.Pow(float64(index_01), 2))

// // 				// Increment the count if the result already exists in the map
// // 				resultFrequency[result]++
// // 			}
// // 		}
// // 	}

// // 	// Convert the resultFrequency map to a slice of [result, frequency] pairs
// // 	type pair struct {
// // 		result    int
// // 		frequency int
// // 	}
// // 	var sortedResult []pair
// // 	for result, frequency := range resultFrequency {
// // 		sortedResult = append(sortedResult, pair{result, frequency})
// // 	}

// // 	// Sort the slice by frequency in descending order
// // 	sort.Slice(sortedResult, func(i, j int) bool {
// // 		return sortedResult[i].frequency > sortedResult[j].frequency
// // 	})

// // 	// Output the frequency of each result
// // 	fmt.Println("\nFrequency of each result:")
// // 	for result, frequency := range resultFrequency {
// // 		fmt.Printf("Result %d appears %d times\n", result, frequency)
// // 	}

// // 	// Output the sorted frequencies in ascending order
// // 	fmt.Println("\nResults sorted from low to high:")
// // 	sort.Slice(sortedResult, func(i, j int) bool {
// // 		return sortedResult[i].frequency < sortedResult[j].frequency
// // 	})
// // 	for _, p := range sortedResult {
// // 		fmt.Printf("Result %d appears %d times\n", p.result, p.frequency)
// // 	}

// // 	// Output the sorted frequencies in descending order
// // 	fmt.Println("\nResults sorted from high to low:")
// // 	sort.Slice(sortedResult, func(i, j int) bool {
// // 		return sortedResult[i].frequency > sortedResult[j].frequency
// // 	})
// // 	for _, p := range sortedResult {
// // 		fmt.Printf("Result %d appears %d times\n", p.result, p.frequency)
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"math"
// 	"sort"
// )

// func main() {
// 	// Creating a map to store the frequency of each result
// 	resultFrequency := make(map[int]int)

// 	// Nested loops to calculate the results and store their frequency
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 7; j++ {
// 			for k := 0; k < 4; k++ {
// 				result := int(math.Pow(float64(i), 2) * math.Pow(float64(j), 2) * math.Pow(float64(k), 2))

// 				// Incrementing the count
// 				resultFrequency[result]++
// 			}
// 		}
// 	}
// }

// // Converting the resultfrequency map to a slice of [result, frequency] pairs
// type pair struct {
// 	result    int
// 	frequency int
// }

// var sortedResult []pair
// for result, frequency := range resultFrequency {
// 	sortedResult = append(sortedResult, pair{result, frequency})
// }

// // Sorting the slice by frequency in descending order
// // Sorting the slice by frequency in descending order//+
// sort.Slice(sortedResult, func(i, j int) bool {
// 	return sortedResult[i].frequency > sortedResult[j].frequency
// })

// // Output the frequency of each result
// fmt.Println("\nFrequency of each result:")
// for result, frequency := range resultFrequency {
// 	fmt.Printf("Result %d appears %d times\n", result, frequency)
// }

// //   git config --global user.email "you@example.com"
// //   git config --global user.name "Your Name"
