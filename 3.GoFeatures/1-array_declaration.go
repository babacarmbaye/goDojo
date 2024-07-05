package main

import "fmt"

func main() {
	// Declaaration without initialization
	var arr1 [5]int
	fmt.Println("Array declared without initialization: ", arr1)

	// Declaration with initialization
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array declared with initialization: ", arr2)

	// Short hand declaration
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with short hand declaration: ", arr3)

	// Initialization with specific elements
	arr4 := [5]int{1: 10, 3: 30}
	fmt.Println("Array initilizing with specific elements: ", arr4)

	// Using ... operator to infer length
	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Array using ... operator to infer length: ", arr5)

	// Array of Arrays (Multi-dimensional Array)
	var multiArray [2][3]int
	multiArray[0] = [3]int{1, 2, 3}
	multiArray[1] = [3]int{4, 5, 6}
	fmt.Println("Multi-dimensional array: ", multiArray)

	// Initializing Multi-dimensional array
	multiArray2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Array multi-dimmensional initialized: ", multiArray2)

	// Array of pointers
	var pointerArray [3]*int
	num1, num2, num3 := 13, 14, 15
	pointerArray[0] = &num1
	pointerArray[1] = &num2
	pointerArray[2] = &num3
	fmt.Println("Array with pointers:")
	for i := 0; i < len(pointerArray); i++ {
		fmt.Printf("pointerArray[%d] = %d\n", i, *pointerArray[i])
	}

	// Array of struct
	type Person struct {
		Name string
		Age  int
	}
	var structArray [2]Person
	structArray[0] = Person{"Aliou", 30}
	structArray[1] = Person{"Fatou", 25}
	fmt.Println("Array of struct :", structArray)
}
