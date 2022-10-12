package main

import (
	"fmt"
)

func main() {
	var sumTo int
	var sumValue int = 0

	fmt.Println("What number are we summing to?")
	fmt.Scanf("%d", &sumTo)
	fmt.Println("You entered: ", sumTo)

	for i := 1; i <= sumTo; i++ {
		sumValue += i
		fmt.Println(sumValue)
	}
}
