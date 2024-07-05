package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// Embedding interfaces
type Animal interface {
	Speaker
	Move() string
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says Tweet!"
}

func (b Bird) Move() string {
	return b.Name + " flies!"
}

// Empty interface
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

// Interface with multiple methods
type Vehicule interface {
	Start()
	Stop()
}

type Car struct {
	Modele string
}

func (c Car) Start() string {
	return c.Modele + " started"
}

func (c Car) Stop() string {
	return c.Modele + " stoped"
}

// Functional interface (Single method interface)
// Simular to Speaker interface

//Interface as a constraint (Generics)
func Describe[T Speaker](t T) {
	fmt.Println(t.Speak())
}

func main() {
	fmt.Println("Interface declaration")
	dog := Dog{Name: "Buddy"}
	bird := Bird{Name: "Tweey"}
	car := Car{Modele: "Tesla"}

	fmt.Println(dog.Speak())
	fmt.Println(bird.Speak())
	fmt.Println(bird.Move())

	PrintAnything("Printing anything, an empty interface can hold anything")
	PrintAnything(46)

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	Describe(dog)
	Describe(bird)
}
