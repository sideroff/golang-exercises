package main

import (
	"fmt"
	"runtime"
)

func main() {
	// by default gomaxprocs = # real threads
	runtime.GOMAXPROCS(100) // os threads
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))

}
