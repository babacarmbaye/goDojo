package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	verbose := flag.Bool("v", false, "Handle verbose mode")
	name := flag.String("name", "World", "Name to greet")
	// Parse flags
	flag.Parse()
	// Use flags
	if *verbose {
		fmt.Println("Verbose mode enbaled !")
	}
	fmt.Printf("Hello %s \n", *name)
}
