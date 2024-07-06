package main

import (
	"fmt"
	"os"
)

func main() {
	PATH := os.Getenv("DATA_PATH")
	data, err := os.ReadFile(PATH)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:\n", string(data))
}
