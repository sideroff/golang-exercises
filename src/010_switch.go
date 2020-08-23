package main

import (
	"fmt"
)

func main() {
	// break is implicit
	switch city := "Pleven"; city {
	case "Varna":
		fmt.Println("Varna")
	case "Sofia", "Pleven":
		fmt.Println("Pleven or Sofia")
	default:
		fmt.Println("default")
	}

	// if you want to fall through, you can explicitly state
	// but keep in mind fallthrough is not checking the next case
	i := 10
	switch {
	case i > 100 && true:
		fmt.Println("Greater than 100")
	case i < 1000:
		fmt.Println("Less than 1000")
		fallthrough
	case i > 1000: // <- this is not evaluated
		fmt.Println("Greater than 1000")
	}

	var a interface{} = 1 //interface = *

	// type switch
	switch a.(type) {
	case int:
		fmt.Println("int")
		break // breaking out of a case
		fmt.Println("this will not execute")
	case float32:
		fmt.Println("float32")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown type")
	}
}
