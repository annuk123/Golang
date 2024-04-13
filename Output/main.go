package main

import "fmt"

func main() {

	//there are three type of Function to output text
	//Print()
	//Println()
	//Printf()

	//THE PRINT() FUNCTION
	//the Print() function prints its arguments with their default format.

	var i, j string = "hello", "World"
	fmt.Print(i)
	fmt.Print(j)
	//output helloWorld

	//if we want to print the arguments in lines, we need to use \n. \n creates new lines.

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	//if we want space between string arguments, we use " "
	fmt.Print(i, " ", j)

	//Print() inserts a space between the arguments if neither are strings

	var x, y = 10, 20
	fmt.Print(x, y)

	//Println() Function
	//the Println() function is similar to Print() with difference that a whitespace is added between the arguments, the arguments,and a newline added at the end:

	var k, l string = "Hello", "World"
	fmt.Println(k, l)

	//Printf() function first formats its argument based on the given formatting verb and then prints them.

	//we will use two formatting verbs
	//%v is used to print the arguments
	//%T is used to print the type of the arguments
	var p string = "Hello"
	var q int = 15
	fmt.Printf("p has value: %v and type: %T\n", p, p)
	fmt.Printf("q has value : %v and type: %T\n", q, q)

}
