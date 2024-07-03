package main

import (
	"fmt"
	"golangBasics/shapes"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Package:")
	radius := 5.0
	areaCircle := shapes.AreaOfCircle(radius)
	fmt.Println("areaCircle =", areaCircle)

	// diameterOfCircle := shapes.diameterOfCircle(radius)
	// fmt.Println("diameterOfCircle =", diameterOfCircle)

	size := 6.0
	areaSquare := shapes.AreaOfSquare(size)
	fmt.Println("areaSquare =", areaSquare)

	// perimeterOfSquare := shapes.perimeterOfSquare(size)
	// fmt.Println("perimeterOfSquare =", perimeterOfSquare)

	id := uuid.New()
	fmt.Println("Generated ID: ", id)
}
