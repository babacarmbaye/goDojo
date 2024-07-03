package main

import "fmt"

type Employee struct {
	Name     string
	Position string
	Salary   float64
}

func updateSalary(emp *Employee, newSalary float64) {
	emp.Salary = newSalary
}

func main() {
	fmt.Println("Pointers:")
	var s int = 10
	var t *int = &s          // & retourn the memory address of the variable
	fmt.Println("s = ", s)   // prints the value of s
	fmt.Println("t = ", t)   // prints the memory address of s
	fmt.Println("*t = ", *t) // prints the value stored in the memory address of t

	fmt.Println("Real example of pointers:")
	em := Employee{
		Name:     "Fatou",
		Position: "DevOps Engineer",
		Salary:   150000,
	}
	fmt.Println("Before update :", em)
	updateSalary(&em, 200000) // For testing pointer here we can pass the object directly without the & and update the method signature
	fmt.Println("After update :", em)
}
