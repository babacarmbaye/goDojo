package main

import "fmt"

func main() {
	fmt.Println("If else if else:")
	var x = 10

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is exactly 10")
	} else {
		fmt.Println("x is less than 10")
	}

	fmt.Println("If else with initialization:")
	if y := 2; y > 20 {
		fmt.Println("y is greater than 20")
	} else {
		fmt.Println("y is less than 20")
	}
}
