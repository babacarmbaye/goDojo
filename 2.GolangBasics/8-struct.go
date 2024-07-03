package main

import "fmt"

func main() {
	fmt.Println("Struct:")
	type Person struct {
		Name string
		Age  int
	}
	var p Person = Person{"Aliou", 30}
	fmt.Println("p =", p)
	fmt.Println("Name =", p.Name)
	fmt.Println("Age =", p.Age)
}
