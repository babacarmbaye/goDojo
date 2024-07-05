package main

import "fmt"

func main() {
	// Using var keyword without initialization
	var slice []int
	fmt.Println("Slice 1 (uninitialized): ", slice)

	// Initilialize slice befor using it
	slice1 := []int{1, 2, 3}
	fmt.Println("Slice initialized: ", slice1)

	// Using the make function
	slice2 := make([]int, 3) // Length and capacity are 3
	fmt.Println("Slice 2 befor putting values: ", slice2)
	slice2[0] = 7
	slice2[1] = 8
	slice2[2] = 9
	fmt.Println("Slice 2 after putting values: ", slice2)
	slice3 := make([]int, 3, 5) // Length 3 and capacity is 4
	fmt.Println("Slice 3 befor putting values: ", slice3)
	slice3[0] = 10
	slice3[1] = 11
	slice3[2] = 12
	fmt.Println("Slice 3 after putting values: ", slice3)
	slice3 = append(slice3, 13)
	slice3 = append(slice3, 14)
	fmt.Println("Slice 3 capacity: ", cap(slice3))
	slice3 = append(slice3, 14)
	fmt.Println("Slice 3 after append: ", slice3)
	fmt.Println("Slice 3 capacity: ", cap(slice3)) // The capacity is multiplied by 2 for memory optimization

	// Slice from array
	array := [5]int{10, 11, 12, 13, 14}
	slice4 := array[1:4] // slicing from index 1 to 3 include
	fmt.Println("Slice 4 :", slice4)
}
