package main

import "fmt"

func main() {
	// upperX := "X"
	// lowerX := "x"

	// // Print ASCII/byte values
	// fmt.Printf("Byte value of 'X': %d\n", upperX[0]) // 88
	// fmt.Printf("Byte value of 'x': %d\n", lowerX[0]) // 120

	upperX := byte('X') // 88  (01011000)
	lowerX := byte('x') // 120 (01111000)

	// Show binary representation
	fmt.Printf("'X' in binary: %08b\n", upperX) // 01011000
	fmt.Printf("'x' in binary: %08b\n", lowerX) // 01111000

	// Show bit difference using XOR
	diff := upperX ^ lowerX
	fmt.Printf("Bit difference (XOR): %08b\n", diff) // 00100000

	// Count different bits
	differentBits := 0
	for i := 0; i < 8; i++ {
		if (upperX>>i)&1 != (lowerX>>i)&1 {
			differentBits++
		}
	}
	fmt.Printf("Number of different bits: %d\n", differentBits) // 1
}
