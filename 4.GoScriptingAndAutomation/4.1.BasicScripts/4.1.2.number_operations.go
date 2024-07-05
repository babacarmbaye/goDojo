package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Basic Arithmetic Operations
	var x, y int = 15, 10
	fmt.Println("Addition: ", x+y)
	fmt.Println("Subtraction: ", x-y)
	fmt.Println("Multiplication: ", x*y)
	fmt.Println("Division: ", x/y)
	fmt.Println("Modulus: ", x%y)

	// Working with Floats
	var m, n float64 = 1.7, 8.9
	fmt.Println("Float Addition: ", m+n)
	fmt.Println("Float Subtraction: ", m-n)
	fmt.Println("Float Multiplication: ", m*n)
	fmt.Println("Float Division: ", m/n)
	// fmt.Println("Float Modulus: ", m%n) modulus not defined on a float variable

	// Type Conversion
	var o int = 13
	var p float64 = float64(o)
	fmt.Println("Integer to float64 :", p)

	q := 3.7
	r := int(q)
	fmt.Println("Float to integer :", r)

	// Working with complex numbers
	var complex1 complex128 = complex(3, 4) // 3 + 4i
	var complex2 complex128 = complex(1, 2) // 1 + 2i
	fmt.Println("Complex Addition:", complex1+complex2)
	fmt.Println("Complex Subtraction:", complex1-complex2)
	fmt.Println("Complex Multiplication:", complex1*complex2)
	fmt.Println("Complex Division:", complex1/complex2)

	// Math functions
	fmt.Println("Power:", math.Pow(2, 3))      // 2^3
	fmt.Println("Square Root:", math.Sqrt(16)) // √16
	fmt.Println("Sin:", math.Sin(math.Pi/2))   // sin(π/2)
	fmt.Println("Cos:", math.Cos(math.Pi))     // cos(π)
	fmt.Println("Log:", math.Log(math.E))      // ln(e)

	// String to Number Parsing
	numStr := "3.14159"
	parsedFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Error parsing string to float:", err)
	} else {
		fmt.Println("Parsed Float:", parsedFloat)
	}
}
