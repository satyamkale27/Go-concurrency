package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()
	go func() {
		anotherChannel <- "cow"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case mgFromAnotherChannel := <-anotherChannel:
		fmt.Println(mgFromAnotherChannel)

	}
}
