package main

import "fmt"

func main() {
	const Pi = 3.14
	//Pi = 400 // enabling this line will throw an error cannot assign to Pi, of cause Pi is const
	fmt.Println("Pi = ", Pi)
	//fmt.Println("xx = ", xx) // enabling this line will throw an undefined error
}
