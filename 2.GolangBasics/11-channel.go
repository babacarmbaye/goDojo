package main

import (
	"fmt"
	"time"
)

// Order represents a customer order
type Order struct {
	ID     int
	Status string
}

// Simulate order processing and sends status through a channel
func processOrder(orderID int, statusChannel chan Order) {
	// simulate some processing time
	time.Sleep(time.Second * 2)

	// sending processed order status back through the channel
	statusChannel <- Order{ID: orderID, Status: "Completed"}
}

func main() {
	fmt.Println("Channel:")
	n := make(chan int)
	fmt.Println("n =", n)

	fmt.Println("Channel real example:")
	// creating a channel for communicating order statuses
	statusChannel := make(chan Order)

	// starting goroutine for processing orders
	for i := 1; i <= 5; i++ {
		fmt.Println("------------------")
		go processOrder(i, statusChannel)
	}

	// receiving and printing order statuses
	for i := 1; i <= 5; i++ {
		fmt.Println("++++++++++++++++++")
		processedOrder := <-statusChannel
		fmt.Printf("Order %d processed: %s\n", processedOrder.ID, processedOrder.Status)
	}
}
