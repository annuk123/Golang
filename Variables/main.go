package main

import "fmt"

//numberOfUser := 30000 // not allowed because it is scope function
const LoginToken string = "Annnuodj" // it is public because the very first letter of variable is in capital letter

//var can be used inside and outside of functions and Variable declaration and value assignment can be done separately

// := Can only be used inside functions and Variable declaration and value assignment cannot be done separately (must be done in the same line)

func main() {
	// fmt.Println("Variables")
	var username string = "Annu"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username) // it is use for know what is type of variable.

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45544356432
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float32 = 255.45544356432
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	//implicit type
	var website = "annu.com"
	fmt.Println(website)

	//no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	//Go Multiple Variable Declaration

	//In Go, it is possible to declare multiple variables in the same line.

	// If you use the type keyword, it is only possible to declare one type of variable per line. as like this example

	var a, b, c, d int = 1, 2, 5, 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//If the type keyword is not specified, you can declare different types of variables in the same line:

	var p, q = 6, "Hello"
	r, s := 7, "World!"

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

	//Go Variable Declaration in a Block

	var (
		x int
		y int    = 1
		z string = "Hello"
	)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
