package main

import "fmt"

type greeter struct {
	name string
}

func (g greeter) greet() { // attach to struct, structs are passed in by value
	fmt.Println(g.name)
}

func (g *greeter) setName(name string) { // you can pass a ref ( eg setter method )
	g.name = name
}

func main() {
	gr := greeter{
		name: "Georgi",
	}

	gr.greet()
	gr.setName("Ivan")
	gr.greet()

}
