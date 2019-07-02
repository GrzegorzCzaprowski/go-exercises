package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ping(channel chan string) {
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(900)))
	fmt.Println(<-channel)
	channel <- "ping"
	wg.Done()
}

func pong(channel chan string) {
	channel <- "pong"
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(900)))
	fmt.Println(<-channel)
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	rand.Seed(time.Now().UnixNano())
	channel := make(chan string)
	for {
		wg.Add(2)
		go ping(channel)
		go pong(channel)
		wg.Wait()
	}
}
