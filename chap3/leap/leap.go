package main

import "fmt"

func main() {
	var year = 2000
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("It's leap year")
	} else {
		fmt.Println("Not a leap year.")
	}
}
