package main

import "fmt"

func main() {
	fmt.Println("Type alias:")
	type ByteSize int64
	var x ByteSize = 1024
	fmt.Println("x =", x)
}
