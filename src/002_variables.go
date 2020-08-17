package main

import (
	"fmt"
	"strconv"
)

var j float64 = 42.5212

// var k := 42. // typing is mandatory on the package level

var actorName string = "Elisabeth Sladen"
var companion string = "Sarah Jane Smith"

var (
	actor2Name string = "Douglas Addams"
	companion2 string = "Arthur Dent"
)

func main() {
	var myVariable int = 24
	myVariable = 42

	i := 52

	var actor2Name = "Elin Pelin" // you can redeclare in inner scopes
	// var actor2Name = "Dora Gabe" // but not in the same scope

	var n float32 = 5.5
	var m int = int(n)

	parsedInt, err := strconv.ParseInt("5.5", 10, 8)
	fmt.Printf("myVariable %v\ni %v\nactorName %v\ncompanion %v\nactor2Name %v\ncompanion2 %v\nn %v\nm %v\nparsedInt %v\nerr %v\n", myVariable, i, actorName, companion, actor2Name, companion2, n, m, parsedInt, err)
}
