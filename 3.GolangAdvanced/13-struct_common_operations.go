package main

import (
	"fmt"
	"reflect"
)

// Person struct definition
type Person struct {
	Name string
	Age  int
}

// Method on the Person struct
func (p *Person) SayHello() {
	fmt.Printf("My name is %s and I am %d years old \n", p.Name, p.Age)
}

// Employee struct with tagged fields
type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Creating an instance of a struct
	elou := Person{Name: "Elou", Age: 39}
	fmt.Println("Instance : ", elou)

	// Accessing and modifying field
	fmt.Println("Instance field befor: ", elou.Name)
	elou.Name = "Falu"
	fmt.Println("Instance field after: ", elou.Name)

	// Pointer struct
	baba := &Person{Name: "Baba", Age: 35}
	fmt.Println("Pointer of struct: ", baba)
	fmt.Println("Pointer of struct: ", *baba)

	// Method on struct
	elou.SayHello()
	baba.SayHello()

	// Tagging struct fields
	emp := Employee{ID: 1, Name: "Buso"}
	t := reflect.TypeOf(emp)
	field, _ := t.FieldByName("Name")
	fmt.Println("Tag on name field: ", field.Tag.Get("json"))
}
