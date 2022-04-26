package main

import (
	"fmt"
)

func main() {

	var name string

	// fmt method

	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
}
