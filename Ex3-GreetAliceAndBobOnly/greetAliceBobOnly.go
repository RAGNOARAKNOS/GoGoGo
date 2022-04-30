package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	goodNames := [...]string{"alice", "bob"}
	var correctName bool = false

	for correctName != true {
		fmt.Println("What is your name?")
		enteredName, _ := reader.ReadString('\n')
		// When reading input on windows we get a new line (\n) and a carriage return (\r) which needs removing
		enteredName = strings.TrimSuffix(enteredName, "\r\n")
		fmt.Println("Hello", enteredName)
		for _, s := range goodNames {
			if s == enteredName {
				correctName = true
			}
		}
	}
	fmt.Println("ACCESS GRANTED - WELCOME ")
}
