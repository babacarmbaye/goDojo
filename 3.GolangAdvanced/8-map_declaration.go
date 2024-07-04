package main

import "fmt"

func main() {
	// Using the var keyword without initialization
	var map1 map[string]int
	fmt.Println("Map declaration with var keyword without init: ", map1)

	// Using a map literal
	map2 := map[string]int{"apple": 5, "banana": 7}
	fmt.Println("Map declared using literal: ", map2)

	// Using make function
	map3 := make(map[string]int)
	fmt.Println("Map declared using make function: ", map3)
	map3["orange"] = 6
	map3["cherry"] = 9
	fmt.Println("Map declared using make function: ", map3) // LIFO (Last In First Out)

	// Using make function with specific size
	map4 := make(map[string]int, 3)
	fmt.Println("Map declared using make function with size: ", map4)
	map4["orange"] = 6
	map4["cherry"] = 9
	map4["banana"] = 19
	fmt.Println("Map declared using make function with specific size: ", map4)

	// Using struct as map value
	type Fruit struct {
		Price float64
		color string
	}
	map5 := make(map[string]Fruit)
	map5["Orange"] = Fruit{10, "orange"}
	map5["Banana"] = Fruit{8, "yellow"}
	fmt.Println("Map declared with struct value: ", map5)

	// Nested Map
	map6 := make(map[string]map[string]int)
	map6["fruit"] = map[string]int{"orange": 5, "banana": 7}
	map6["vegetable"] = map[string]int{"carrot": 3, "peas": 2}
	fmt.Println("Nested map: ", map6)
}
