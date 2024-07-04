package main

import "fmt"

// Function to modify the value of a pointer
func modifyValue(ptr *int) {
	*ptr = 3000
}

func main() {
	// Dereferencing pointer
	a := 20
	ptrA := &a
	fmt.Println("Original value of a :", a)
	fmt.Println("Dereferencing value :", *ptrA)

	// Changing the value at a pointer
	*ptrA = 40
	fmt.Println("After changing the value from pointer :", a)

	// Comparing pointers
	b := 20
	ptrB := &b
	fmt.Println("ptrA and ptrB contains the same address? ", ptrA == ptrB)
	p1 := &a
	p2 := &b
	p3 := &a
	fmt.Println("p1 and p2 contains the same address ? ", p1 == p2)
	fmt.Println("p1 and p3 contains the same address ? ", p1 == p3)

	// Passing pointer to function
	fmt.Println("Value of 'a' befor calling function: ", a)
	modifyValue(ptrA)
	fmt.Println("Value of 'a' after calling function: ", a)
}
