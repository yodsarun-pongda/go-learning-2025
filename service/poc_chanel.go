package service

import (
	"fmt"
	"time"
)

func PocChannel() {
	// Create a channel
	ch := make(chan string)

	// Start a goroutine to send data to the channel
	go func() {
		fmt.Println("Start goroutine...")
		time.Sleep(20 * time.Second)
		fmt.Println("Sending data to channel...")
		ch <- "The job have been succeed!"
	}()

	// Wait for the goroutine to finish
	time.Sleep(10 * time.Millisecond)
	fmt.Println("waiting for channel...")
	message := <-ch
	fmt.Println("Got message from Job:", message)
}
