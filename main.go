package main

import (
	receiver "chatApplication/recevier"
	"chatApplication/sender"
	"fmt"
)

func main() {
	// Create a channel for communication between sender and receiver
	ch := make(chan string)

	go sender.Send(ch)   // Start the sender in a separate goroutine
	receiver.Receive(ch) // Start the receiver to receive messages from the channel
	fmt.Print("done")
}
