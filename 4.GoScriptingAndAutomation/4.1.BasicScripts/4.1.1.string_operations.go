package main

import (
	"fmt"
	"strings"
)

func main() {
	// String concatenation
	str1 := "Hello"
	str2 := "World"
	concated := str1 + " " + str2
	fmt.Println("Concated string :", concated)

	// Splitting a string
	sentence := "Go is an open source programming langage"
	words := strings.Split(sentence, " ")
	fmt.Println("Words in sentence :", words)

	// Replacing a substring
	replaced := strings.Replace(sentence, "Go", "Golang", 1)
	fmt.Println("Replaced :", replaced)

	// Upper and Lower case
	lower := strings.ToLower(sentence)
	upper := strings.ToUpper(sentence)
	fmt.Println("Lower case: ", lower)
	fmt.Println("Upper case: ", upper)

	// Trimming
	nonTrimed := "    Trim me      "
	trimed := strings.TrimSpace(nonTrimed)
	fmt.Println("Trimed :", trimed)

	// Substring using slicing
	// Note: Go don't have a built-in "substring" method, but this can be acheived using slice
	start := 3
	end := 11
	if start < end && end <= len(sentence) {
		substring := sentence[start:end]
		fmt.Println("Substring: ", substring)
	} else {
		fmt.Println("Invalid range for substring")
	}

	// Checking if a string contains a substring
	contains := strings.Contains(sentence, "source")
	fmt.Println("Contains source :", contains)

	// Finding a substring's 'index'
	index := strings.Index(sentence, "source")
	fmt.Println("Index :", index)
}
