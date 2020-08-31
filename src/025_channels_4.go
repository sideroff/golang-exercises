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
		i := <-ch
		fmt.Println(i)

		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 20
		wg.Done()
	}(ch)

	wg.Wait()

}
