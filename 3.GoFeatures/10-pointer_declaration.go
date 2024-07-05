package main

import "fmt"

func main() {
	// Basic pointer declaration
	var ptr1 *int
	fmt.Println("Using var keyword to declare the pointer: ", ptr1)

	// Pointer from variable
	var val int = 12
	ptr2 := &val
	fmt.Printf("Pointer %d declared from existing variable %d \n ", ptr2, val)

	// Pointer with new
	ptr3 := new(int)
	*ptr3 = 100
	fmt.Println("Pointer with new :", ptr3)
	fmt.Println("Value of pointer with new :", *ptr3)

	// Pointer of pointer
	ptr4 := &ptr3
	fmt.Println("Pointer to pointer: ", ptr4)
	fmt.Println("Value pointer to pointer: ", *ptr4)
	fmt.Println("Value value pointer to pointer: ", **ptr4)
}
