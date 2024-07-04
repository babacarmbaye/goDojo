package main

import "fmt"

func modifySlice(slice []int) {
	if len(slice) > 0 {
		slice[0] = 999
	}
}

func main() {
	// Initializing slice
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Original slice: ", slice)

	modifySlice(slice)

	fmt.Println("After modification slice: ", slice)
}
