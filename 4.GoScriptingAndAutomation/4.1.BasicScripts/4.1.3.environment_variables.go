package main

import (
	"fmt"
	"os"
)

func main() {
	// Settings an env var
	os.Setenv("MY_ENV_VAR", "Hello world")

	// Reading an env var
	myVar := os.Getenv("MY_ENV_VAR")
	if myVar == "" {
		fmt.Println("MY_ENV_VAR is not set")
	} else {
		fmt.Printf("MY_ENV_VAR: %s \n", myVar)
	}

	// Listing all environment variable
	// fmt.Println("All env var: \n")
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }
}
