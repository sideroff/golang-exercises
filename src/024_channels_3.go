package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			ch <- 42 // blocking => deadlock
			wg.Done()
		}()
	}

	wg.Wait()

}
