package main

import "fmt"

// Function to modify an element of an array
func modifyArray(arr [5]int, index, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {
	fmt.Println("Array Common Operations")
	// Initialize an array
	originArray := [5]int{1, 2, 3, 4, 5}

	// Print the length of the array
	fmt.Println("The length of the array is :", len(originArray))

	// Copy the array and print the copied array
	var copiedArray = originArray
	fmt.Println("Copied array: ", copiedArray)

	// Modify the third element (index 2) of the array to 10
	modifiedArray := modifyArray(originArray, 2, 10)
	fmt.Println("Modified array: ", modifiedArray)

	// Compare the original and modified array
	isEqual := originArray == modifiedArray
	fmt.Println("Are original and modified arrays equal ?", isEqual)

	// Compare the original and copied array
	isCopiedEqual := originArray == copiedArray
	fmt.Println("Are original and copied arrays equal ?", isCopiedEqual)

	// Define an array of integers
	numbers := [5]int{10, 20, 30, 40, 50}
	// Using a traditional for loop to iterate over the array
	fmt.Println("Iterating using traditional for loop")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d, Value: %d \n", i, numbers[i])
	}
	// Using a range loop to iterate over the array
	fmt.Println("Iterating using range loop")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d \n", index, value)
	}
}
