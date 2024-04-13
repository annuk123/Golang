package main

import "fmt"

//The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

//The value of a constant must be assigned when you declare it.

//Type Constant
const A int = 1

//Untype Constant
const B = 1

const PI = 3.14

func main() {
	fmt.Println(PI)

	//Constant Rules
	//constant names follow the same naming rules as variables
	//constant names are usually written in Uppercase letter
	//Constant can be declare both inside and outside of a function

	//

	//Constant Types
	//There are two types of constants
	//Typed constants
	// Untyped constants

	//Type Constant
	//Typed constant are declare with a defined type
	fmt.Println(A)

	//Untyped Constants
	// Untyped constant are declared without a type.
	//In this case complier decides the type of constant, based on the value.

	fmt.Println(B)

	//Constants: Unchangeable and Read only
	//When a constant is declared, it is not possible to change the value later.
	//const A  =  1
	//A  =  3
	//fmt.Println(A)

	//Multiple Constant Declaration
	//Multiple constants can be grouped together into a block for readablity

	const (
		X int = 1
		Y     = 3.14
		Z     = "HII!"
	)

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)

}
