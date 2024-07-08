package main

import "fmt"

func main() {
	var input string
	fmt.Println("Enter a value:")
	fmt.Scanln(&input)
	fmt.Println("Printed value: ", input)
}
