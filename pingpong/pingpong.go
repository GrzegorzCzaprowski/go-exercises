package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	rand.Seed(time.Now().UnixNano())

	for {
		go func() {
			time.Sleep(time.Duration(rand.Intn(1000-100)+100) * time.Millisecond)
			channel <- "ping"
		}()
		fmt.Println(<-channel)

		go func() {
			time.Sleep(time.Duration(rand.Intn(1000-100)+100) * time.Millisecond)
			channel <- "pong"
		}()
		fmt.Println(<-channel)
	}
}
