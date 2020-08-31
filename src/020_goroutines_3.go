package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // infinite readers, 1 writer, if writer => no read

func main() {

	runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // read lock
		go sayHello()
		m.Lock() // w lock
		go increment()
	}

	wg.Wait()
}

func sayHello() {
	fmt.Println("hello ", counter) // race condition => counter is unreliable
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock() // w unlock
	wg.Done()
}
