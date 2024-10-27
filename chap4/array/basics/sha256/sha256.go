package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// %x to print all the elements of an array or slice of bytes in hexadecimal,
	// %t to show a boolean,
	// %T to display the type of a value
	fmt.Printf("%x\n%X\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
