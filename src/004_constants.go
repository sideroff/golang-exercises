package main

func main() {

	const myConst = 5
	const myOtherConst int16 = 6
	myConst + myOtherConst // works, compiler basically replaces aliases with real value ( ... => 5 + 6)
	// constants must be imediately evaluated
	// const myOtherConst int;

}
