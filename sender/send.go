package sender

import "fmt"

func Send(ch chan string) {

	// ch := make(chan string)

	msg := "hello from sender to receiver"
	fmt.Println("Sending message:", msg)
	ch <- msg

}
