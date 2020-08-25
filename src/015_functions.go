package main

import "fmt"

func printMsg(msg string, msg2 string) {
	fmt.Println(msg, msg2)
}

func printMsg2(msg, msg2 string) { // you can use coma delimited list for arguments
	fmt.Println(msg, msg2)
}

func printMsg3(msg, msg2 *string) { // you can pass pointers as arguments
	fmt.Println(msg, msg2)
	*msg = "will change parameter"
}

func printMsg4(msgs ...string) { // you can use arguments[]
	fmt.Println(msgs)
}

func returnMsg(msg string) string { // return types
	return msg
}
func sum(values ...int) (result int) { // you can specity a var that will be locally available as a return
	for _, v := range values {
		result += v
	}
	return // implicit return result here
}

func divide(a, b float64) (float64, error) { // you can return tuples (2+ items)
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil // explicitly return nil as a "no error" flag
}
func main() {
	// entry point - package main => func main (void)

	printMsg("Hello world", "!")
	printMsg4("test", "1", "2", "3")

	func() {
		println("annonymous fn ( iife ) ")
	}()

	var f func(float64, float64) (float64, error) = divide // functions are first class citizens

	println(f) // pointer to fn
}
