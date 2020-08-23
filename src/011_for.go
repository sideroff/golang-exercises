package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	j := 0
	fmt.Println()

	// while
	for j < 3 {
		fmt.Println(j)
		j++
	}
	fmt.Println()

	for {
		fmt.Println("infinite loop")
		break
		continue
	}
	fmt.Println()

LoopName:
	for {
		for {
			fmt.Println("break will break closest loop if not labeled")
			break LoopName
		}
	}
	fmt.Println()

	// foreach
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Printf("%v, %v\n", k, v)
	}
}
