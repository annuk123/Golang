package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok|| err

	input, _ := reader.ReadString('\n') // we can write it input, err if you have chance to catch error and _, err also we can write. if you thought error not be found then you can use input, _
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)
}
