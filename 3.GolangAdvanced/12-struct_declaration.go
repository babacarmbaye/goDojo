package main

import "fmt"

func main() {
	// Basic declaration of a struct
	type Person struct {
		Name string
		Age  int
	}

	laye := Person{Name: "Laye", Age: 20}
	fmt.Println("Instancied object: ", laye)

	// Using the new keyword
	abdou := new(Person)
	abdou.Name = "Abdou"
	abdou.Age = 25
	fmt.Println("Using keyword new :", abdou) // variable is pointer with &
	fmt.Println("Using keyword new :", *abdou)

	// Anonymous struct
	anon := struct {
		Country string
		Code    int
	}{
		Country: "FR",
		Code:    33, // here the ',' is mandatory
	}
	fmt.Println("Anonymous struct :", anon)

	// Nested struct
	type Address struct {
		City   string
		Street string
	}
	type Employee struct {
		PersonalInfos Person
		Address       Address
	}
	emp := Employee{
		PersonalInfos: Person{Name: "Fatou", Age: 30},
		Address:       Address{City: "Paris", Street: "22 Test Street"},
	}
	fmt.Println("Nested struct :", emp)
	fmt.Println("Nested struct :", emp.PersonalInfos.Name)
	fmt.Println("Nested struct :", emp.PersonalInfos.Name)

	// Embedded fiels
	type Manager struct {
		Person
		Department string
	}
	mgr := Manager{
		Person:     Person{Name: "Awa", Age: 35},
		Department: "Finance",
	}
	fmt.Println("Embedded fields: ", mgr)
	fmt.Println("Embedded fields: ", mgr.Age) // here we can directly get the nested attribute
	fmt.Println("Embedded fields: ", mgr.Person.Age)
}
