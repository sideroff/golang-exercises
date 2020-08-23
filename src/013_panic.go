package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(1)
	panicker()
	fmt.Println(2)

}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { // calling recover actually captures the exception
			log.Println("Error: ", err)
			// panic(err) // we can repanic if we cannot handle the exception
		}
	}() // defered iife

	panic("we panicked")
	fmt.Println("not gonna execute")

}
