package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var guessedNumber = 10
	var count = 10
	for count > 0 {
		var randomNumber = rand.Intn(10)
		// if randomNumber == guessedNumber {
		// 	fmt.Println("Found")
		// }
		fmt.Println(randomNumber)
		count--
	}
}
