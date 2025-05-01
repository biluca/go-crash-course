package main

import "fmt"

func main() {

	var channel = make(chan int, 7)

	go process(channel)
	for value := range channel {
		fmt.Println(value)
	}

}

func process(channel chan int) {
	defer close(channel)
	for i := 0; i < 7; i++ {
		channel <- i
	}

}
