package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // buffered channel can mean we receive multiple messages, but if we dont read all of them they are lostddd
	wg.Add(2)

	// <- operator is blocking

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // channels can be closed

		// ch <- 2   // panic - sending a msg on a closed channel

		wg.Done()
	}(ch)

	wg.Wait()

}
