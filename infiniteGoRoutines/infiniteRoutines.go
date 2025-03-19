package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done: // if we received done already then return
			return
		default:
			fmt.Println("Doing work")
		}
	}

}

func main() {
	done := make(chan bool) // channel created with make
	go doWork(done)

	time.Sleep(time.Second * 3) // this allows run goroutine 3 secs
	close(done)                 // closes goroutine of child dowork
}
