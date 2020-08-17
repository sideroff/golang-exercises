package main

import "fmt"

func main() {

	// bools
	var n bool = true
	t := 1 == 1
	f := 1 == 2

	var def bool // types have default vals, false in this case

	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", t, t)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", def, def)

	//int

}
