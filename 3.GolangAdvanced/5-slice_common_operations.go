package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initializing slice
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	fmt.Println("Original slice :", slice)

	// Print length and capacity
	fmt.Printf("length: %d, capacity: %d \n", len(slice), cap((slice)))

	// Appending to a slice
	slice = append(slice, 7)
	fmt.Println("After appending: ", slice)

	// Slicing a slice
	subSlice := slice[2:5]
	fmt.Println("Sub slice: ", subSlice)

	// Delete an element from slice (e.g at index 3)
	slice = append(slice[:3], slice[4:]...)
	fmt.Println("After deleting an element: ", slice)

	// Check if a slice is empty
	if len(slice) == 0 {
		fmt.Println("Slice is empty")
	} else {
		fmt.Println("Slice is not empty")
	}

	// Sort slice
	sort.Ints(slice)
	fmt.Println("Sorted slice :", slice)

	// Find an element in slice
	elementToFind := 5
	isFound := false
	for _, item := range slice {
		if item == elementToFind {
			isFound = true
			break
		}
	}
	if isFound {
		fmt.Printf("Element %d was found in the slice %d \n", elementToFind, slice)
	} else {
		fmt.Printf("Element %d was not found in the slice %d \n", elementToFind, slice)
	}
}
