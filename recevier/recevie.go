package receiver

func Receive(ch chan string) {
	msg := <-ch
	println("Received message:", msg)
}
