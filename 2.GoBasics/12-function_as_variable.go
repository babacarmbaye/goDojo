package main

import "fmt"

func main() {
	fmt.Println("Function as variable:")
	var v func(int) int
	v = func(i int) int { return i * i }
	result := v(5)

	fmt.Println("result = ", result)
}
