package main

import (
	"fmt"
	"time"
)

func printCount(label string, count int, ch chan int) {
	for i := 0; i < count; i++ {
		ch <- i
		fmt.Println(label, i)
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go printCount("Goroutine", 5, ch)

	for value := range ch {
		fmt.Println("Main received :", value)
	}
}
