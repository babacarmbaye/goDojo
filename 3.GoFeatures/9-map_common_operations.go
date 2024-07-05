package main

import "fmt"

func main() {
	// Initialize map
	fruitsMap := map[string]int{"apple": 5, "banana": 4}

	// Check length
	fmt.Println("The length of the map is: ", len(fruitsMap))

	// Adding element
	fruitsMap["orange"] = 8
	fmt.Println("Map after element added: ", fruitsMap)

	// Retreiving an element
	applePrice := fruitsMap["apple"]
	fmt.Println("Price of apple :", applePrice)

	// Checking existence
	price, exists := fruitsMap["kiwi"]
	if exists {
		fmt.Println("Price of kiwi is: ", price)
	} else {
		fmt.Println("Kiwi does not exist in the map")
	}
	orangePrice, exists := fruitsMap["orange"]
	if exists {
		fmt.Println("Price of orange is: ", orangePrice)
	} else {
		fmt.Println("Orange does not exist in the map")
	}

	// Iterating on map
	fmt.Println("Fruit : Price")
	for key, value := range fruitsMap {
		fmt.Printf("%s : %d \n", key, value)
	}

}
