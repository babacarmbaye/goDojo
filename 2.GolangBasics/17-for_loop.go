package main

import "fmt"

func main() {
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("For loop as a while loop:")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Infinite for loop:")
	j := 0
	for {
		// some code here
		if j > 100 {
			fmt.Println("breaking...")
			break // break out of the loop at this point
		}
		fmt.Println(j)
		j++
	}

	fmt.Println("for Loop with range (for slices, arrays, maps, strings):")
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, value %d \n", index, value)
	}

	fmt.Println("for Loop with range (for strings):")
	for index, runeValue := range "hello" {
		fmt.Printf("Index: %d, Rune %c \n", index, runeValue)
	}
}
