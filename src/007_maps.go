package main

import "fmt"

func main() {

	dict := make(map[string]int)

	dict = map[string]int{
		"Varna":   1,
		"Burgas":  2,
		"Plovdiv": 3,
		"Sofia":   4,
	}
	dict["Pleven"] = 5

	fmt.Printf("%v\n", dict) // unordered because of hashes
	fmt.Printf("%v\n", dict["Sofia"])
	fmt.Printf("%v\n", dict["Gosho"]) // missing keys are = to default value

	gosho, ok := dict["Gosho"]        // "coma ok" syntax
	fmt.Printf("%v, %v\n", gosho, ok) // ok flag shows if key was found or not

	delete(dict, "Varna")    // delete
	delete(dict, "Gosho")    // delete of missing key just passes
	fmt.Printf("%v\n", dict) // unordered because of hashes

}
