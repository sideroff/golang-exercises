package main

import (
	"fmt"
	"reflect"
)

// Doctor <-required; comments are convention
type Doctor struct {
	number      int
	actorName   string
	companions  []string
	PublicField string
}

// Animal <docs>
// backticks string is called a tag and is meaningless by itself
// its rather some "metadata" provided to any interested framework via reflection
// ex validation
type Animal struct {
	name string `required min: "3"`
}

// Bird <docs>
type Bird struct {
	Animal
	canFly bool
}

func main() {

	aDoctor := Doctor{
		number:     1,
		actorName:  "David Tannant",
		companions: []string{"Rose", "Martha"},
	}

	// structs are value types

	fmt.Printf("%v\n", aDoctor)
	fmt.Printf("%v\n", aDoctor.actorName)
	fmt.Printf("%v\n", aDoctor.companions[0])

	aBird := Bird{}
	aBird.name = "Pe"
	aBird.canFly = false
	fmt.Printf("%v\n", aBird)

	fmt.Printf("%v\n", aBird.Animal.name)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Printf("tag= %v\n", field.Tag)

}
