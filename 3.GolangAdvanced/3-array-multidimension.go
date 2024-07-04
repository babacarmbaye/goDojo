package main

import "fmt"

func main() {
	// Initialize the array
	var matrix [3][3]int
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = count
			count++
		}
	}

	// Print array
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
