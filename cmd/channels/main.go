package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 1)

	go fetchData(messages)

	for {
		message, ok := <-messages

		if !ok {
			break
		}

		fmt.Println(message)
	}
}

func fetchData(messages chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		messages <- i
	}
}
