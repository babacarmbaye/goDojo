package main

import "fmt"

func main() {
	// Initializing slice
	slice := []string{"apple", "banana", "cherry", "date", "orange"}

	// Iterating over the slice using for loop
	fmt.Println("Iterating with for loop")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d, value: %s \n", i, slice[i])
	}

	// Iterating over the slice using range
	fmt.Println("\nIterating with range")
	for i, item := range slice {
		fmt.Printf("Index: %d, value: %s \n", i, item)
	}

}
