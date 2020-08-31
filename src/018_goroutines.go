package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	// go sayHello()

	var msg = "closure and annonymous fn"
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "reasigned"

	wg.Add(1)
	// pass data in goroutine
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "reasigned2"

	wg.Wait()
	// time.Sleep(100 * time.Microsecond)
}

func sayHello() {
	fmt.Println("hello")
}
