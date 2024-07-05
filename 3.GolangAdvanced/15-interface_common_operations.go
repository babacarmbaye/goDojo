package main

import "fmt"

// Interface declaration
type Greeter interface {
	Greet() string
}

type Person2 struct {
	Name string
}

func (p Person2) Greet() string {
	return "Hello, " + p.Name
}

// Type assertion
func PrintDetails(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Println("It's a string: ", str)
	} else {
		fmt.Println("Not a string")
	}
}

// Interface with type switches
func Describe2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("String: ", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	// Implementing an interface
	person := Person2{Name: "Aliou"}
	var greeter Greeter = person
	fmt.Println(greeter.Greet())

	// Type Assertions
	PrintDetails("Hello")
	PrintDetails(42)

	// Interface with type switches
	Describe2("Hello, World")
	Describe2(2024)
	Describe2(3.14)
}
