package main

import "fmt"

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(1981))
}
