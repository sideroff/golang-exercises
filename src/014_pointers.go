package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 1
	var b *int = &a

	// &<var> = address of var
	// *<pointer> = dereference, get value of pointer

	fmt.Printf("%v, %v\n", a, b)  // a = 1, b = <memory address>
	fmt.Printf("%v, %v\n", a, *b) // a = 1, b = *<memory address> = 1

	a = 2

	fmt.Printf("%v, %v\n", a, *b) // a = 2, b = *<memory address> = 2

	*b = 3

	fmt.Printf("%v, %v\n", a, *b) // a = 3, b = *<memory address> = 3

	// -----

	var c *int

	fmt.Printf("%v\n", c) // c = <nil>; default pointer value

	var d *myStruct
	d = new(myStruct)
	(*d).foo = 1 // . operator takes precedence over * fml

	fmt.Printf("%v\n", (*d).foo) // (*d).foo = 1

	var e *myStruct
	e = new(myStruct)
	e.foo = 1 // (*e).foo = e.foo ( implied, syntax sugar)

	fmt.Printf("%v\n", e.foo)
}
