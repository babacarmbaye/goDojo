package main

import (
	"errors"
	"fmt"
)

// Custom Error
type CustomError struct {
	msg string
}

// CustomError implements the error interface
func (e *CustomError) Error() string {
	return e.msg
}

// SomeFunction demonstrates returning a custom error
func SomeFunction(flag bool) error {
	if !flag {
		return &CustomError{"Custom error occured"}
	}
	return nil
}

// SafeFunction demonstrates using panic and recover
func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic: ", r)
		}
	}()
	// This would cause a panic
	panic("A problem occured")
}

// Divide demonstrates division by zero error
func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Division by zero")
	}
	return x / y, nil
}

func main() {
	// Simple error returning
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	// Example of custom error
	err2 := SomeFunction(false)
	if err2 != nil {
		fmt.Println("Error :", err2)
	}

	// Example of panic and recover
	SafeFunction()
}
