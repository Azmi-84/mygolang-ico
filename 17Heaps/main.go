package main

import "fmt"

// MaxHeap represents a Max Heap data structure
type MaxHeap struct {
	arr []int
}

// Insert adds a new element to the heap and maintains the heap property
func (hp *MaxHeap) Insert(key int) {
	hp.arr = append(hp.arr, key)
	hp.heapifyMaxUp(len(hp.arr) - 1)
}

// heapifyMaxUp maintains the max-heap property after insertion
func (hp *MaxHeap) heapifyMaxUp(index int) {
	for index > 0 && hp.arr[parent(index)] < hp.arr[index] {
		hp.swap(parent(index), index)
		index = parent(index)
	}
}

// parent returns the index of the parent node
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the left child
func left(i int) int {
	return 2*i + 1
}

// right returns the index of the right child
func right(i int) int {
	return 2*i + 2
}

// swap swaps the values at two indices
func (hp *MaxHeap) swap(int1, int2 int) {
	hp.arr[int1], hp.arr[int2] = hp.arr[int2], hp.arr[int1]
}

func main() {
	m := &MaxHeap{} // Initialize the MaxHeap instance
	fmt.Println("MaxHeap before insertion:", m.arr)

	makeHeap := []int{1, 2, 3} // Correct slice declaration
	for _, v := range makeHeap {
		m.Insert(v) // Insert each value into the heap
	}
	fmt.Println("MaxHeap after insertion:", m.arr) // Output the heap after insertions
}
