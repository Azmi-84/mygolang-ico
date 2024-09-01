package main

import (
	"fmt"
	"sort"
)

func main() {
	// var Slice = []string{"a", "b", "c", "d", "e"}
	// fmt.Println(Slice)
	// fmt.Printf("Type of slice is %T\n", Slice)

	// Adding element to slice
	// Slice = append(Slice, "f")
	// fmt.Println(Slice)

	// Slicing the slice
	// Slice_00 := Slice[1:]
	// fmt.Println(Slice_00)

	// slice_01 := Slice[:3]
	// fmt.Println(slice_01)

	slice_02 := make([]int, 5)

	slice_02[0] = 10
	slice_02[1] = 2
	slice_02[2] = 5
	slice_02[3] = 4
	slice_02[4] = 3
	// slice_02[5] = 6 // this will give an error because the size of the slice is 5

	// fmt.Println(slice_02)

	// if we want to add more elements to the slice_02 we can use append function
	slice_02 = append(slice_02, 6, 9, 8, 7, 1)
	fmt.Println(slice_02)

	// Sorting the slice
	sort.Ints(slice_02)
	fmt.Println(slice_02)
	// fmt.Println(sort.IntsAreSorted(slice_02))

	// Reversing the slice
	for i, j := 0, len(slice_02)-1; i < j; i, j = i+1, j-1 {
		slice_02[i], slice_02[j] = slice_02[j], slice_02[i]
	}
	fmt.Println(slice_02)

	// Copying the slice but the size of the slice should be the same as the original slice
	// slice_03 := make([]int, 5)
	// copy(slice_03, slice_02)
	// fmt.Println(slice_03)

	// Removing element from the slice
	// slice_04 := append(slice_02[:2], slice_02[3:]...)
	// fmt.Println(slice_04)

}
