package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = append(slice, 10)

	fmt.Printf("Values: %v\n", slice)
	fmt.Printf("Length: %v\n", len(slice))
	fmt.Printf("Capacity: %v\n\n", cap(slice)) // double ( underlying array doubles in size )

	a := slice[:]   // all elements
	b := slice[3:]  // all elements from 3rd index onward inclusive
	c := slice[:6]  // all elements until the 6rd index exclusive
	d := slice[3:6] // all elements from the 3rd to the 6th index
	d[0] = 15       // slice ranges are still referencing the same value
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)

	copy := slice
	copy[0] = 5

	fmt.Printf("Values: %v\n", slice)

	slice = append(slice, []int{11, 12}...) // spread operator

	e := []int{1, 2, 3, 4, 5}
	f := append(e[:2], e[3:]...)
	fmt.Printf("e: %v\n", e) // e is changed here
	fmt.Printf("f: %v\n", f)

}
