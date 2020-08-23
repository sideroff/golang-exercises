package main

import "fmt"

func main() {
	fmt.Println(1)
	// anything defered is added to a stack (LIFO) that is executed once the function returns
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)

	test := "fizz"
	defer fmt.Println(test) // has some sort of closure, value here is fizz regardless of redefinition
	test = "buzz"

}
