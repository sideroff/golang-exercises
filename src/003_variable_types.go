package main

import "fmt"

func main() {

	// bools
	var n bool = true
	t := 1 == 1

	var def bool // types have default vals, false in this case

	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", t, t)
	fmt.Printf("%v %T\n", def, def)

	//int ( bits dependant on the os, always >=32 bits)
	var signedInt int = 32
	fmt.Printf("%v %T\n", signedInt, signedInt)

	var unSignedInt uint = 32
	fmt.Printf("%v %T\n", unSignedInt, unSignedInt)

	// type conversion does not allow us to merge types
	// signedInt + unSignedInt



}
