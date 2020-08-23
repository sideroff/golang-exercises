package main

import (
	"fmt"
)

func main() {

	if true && true || false {
		fmt.Println("if example")
	}

	dict := map[string]int{
		"Sofia":   1,
		"Varna":   2,
		"Plovdiv": 3,
	}

	if sofia, ok := dict["Sofia"]; ok {
		fmt.Printf("if with initialized var sofia= %v", sofia)
	}
}
