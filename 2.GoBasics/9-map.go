package main

import "fmt"

func main() {
	fmt.Println("Map:")
	m := map[string]int{"Aliou": 30, "Samba": 25}
	fmt.Println("m =", m)
	fmt.Println("Aliou =", m["Aliou"])
	fmt.Println("Samba =", m["Samba"])
}
