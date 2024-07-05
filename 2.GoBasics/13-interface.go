package main

import "fmt"

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	fmt.Println("interface:")
	var g Geometry = Rect{4, 5}

	fmt.Println("object =", g)
	fmt.Println("Area   =", g.Area())
	fmt.Println("Perimeter =", g.Perimeter())
}
