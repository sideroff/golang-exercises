package main

import (
	"fmt"
)

func main() {

	var students [3]string

	students[0] = "Ivan"
	students[1] = "Gosho"
	students[2] = "Pesho"

	fmt.Printf("%v\n", len(students))

	var matrix [3][3]int

	matrix[0] = [3]int{0, 0, 0}
	matrix[1] = [3]int{1, 1, 1}
	matrix[2] = [3]int{2, 2, 2}

	fmt.Println("\narrays are value types by default")
	a := [...]int{1, 2}
	b := a

	b[0] = 5

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("\nbut can be forced to use references with a pointer")
	c := [...]int{1, 2}
	d := &c

	d[0] = 5

	fmt.Println(c)
	fmt.Println(d)
}
