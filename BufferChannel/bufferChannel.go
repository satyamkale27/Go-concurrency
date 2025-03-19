package main

import "fmt"

func main() {
	charChannel := make(chan string, 3)

	char := []string{"a", "b", "c"}
	for _, s := range char {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)
	for result := range charChannel {
		fmt.Println(result)
	}
}
