package main

import "fmt"

func main() {
	fmt.Println("Switch:")
	var day = 22

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("invalid day")
	}

	var x = -2
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is positive")
	}
}
