package main

import (
	"crypto/sha256"
	"fmt"
)

// Function to count the number of different bits in two SHA-256 hashes
func countDiffBits(hash1, hash2 [32]byte) int {
	diffBits := 0

	// Compare each byte in the two hashes
	for i := 0; i < len(hash1); i++ {
		// XOR the bytes and count the bits that are 1
		diff := hash1[i] ^ hash2[i]
		diffBits += countBits(diff)
	}

	return diffBits
}

// Helper function to count the number of 1 bits in a byte
func countBits(b byte) int {
	count := 0
	for b != 0 {
		count += int(b & 1)
		b >>= 1
	}
	return count
}

func main() {
	data1 := "x"
	data2 := "X"

	// Compute SHA-256 hashes
	// hash1 and hash2 both are [32]uint8 types
	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))

	fmt.Printf("%T\n%T\n", hash1, hash2)

	// Count the different bits
	diffBits := countDiffBits(hash1, hash2)
	fmt.Printf("Number of differing bits: %d\n", diffBits)
}
