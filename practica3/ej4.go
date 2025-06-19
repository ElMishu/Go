package main

import (
	"fmt"
)

func pxng(m chan string, str string, wait chan bool) {
	wait <- true
	m <- str
}

func main() {
	messages := make(chan string)
	wait := make(chan bool)
	for i := 0; i < 5; i++ {
		go pxng(messages, "PING", wait)
		<-wait // Espera a que se consuma el mensaje PING
		go pxng(messages, "PONG", wait)
		<-wait // Espera a que se consuma el mensaje PONG
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-messages)
	}
}
