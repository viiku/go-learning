// Enhance comma so that it deals correctly with floating-point numbers and an
// optional sign

package main

import "fmt"

func main() {
	s := "abc"
	// abc
	fmt.Println(s)
	b := []byte(s)
	// [97,98,99]
	fmt.Println(b)
	s2 := string(b)
	// abc
	fmt.Println(s2)
}

// case 1.
// the []byte(s) conversion allocates a new byte array holding a copy of the bytes
// of s, and yields a slice that references the entirety of that array.
// In general copying is required to
// ensure that the bytes of s remain unchanged even if those of b are subsequently modified.

// case 2.
// The conversion from byte slice back to string with string(b) also makes a copy, to ensure
// immutability of the resulting string s2
