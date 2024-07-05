package main

import "fmt"

func main() {
	fmt.Println("array:")
	var arr [5]int = [5]int{1, 2, 3, 4, 5} // The size of array's content cannot change
	fmt.Println("arr =", arr)
}
