// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.

package main

import (
	"fmt"
	"strings"
)

func isAnagram(str1, str2 string) bool {
	// If lengths are different, they can't be anagrams
	if len(str1) != len(str2) {
		return false
	}

	// Convert to lowercase to make comparison case-insensitive
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Create a map to store character counts
	charCount := make(map[rune]int)

	// Count characters in first string
	for _, char := range str1 {
		charCount[char]++
		fmt.Println(charCount[char])
	}

	// Subtract counts for second string
	for _, char := range str2 {
		charCount[char]--
		// If character count becomes negative, strings aren't anagrams
		if charCount[char] < 0 {
			return false
		}
	}

	// Check if all counts are zero
	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	result := isAnagram(s1, s2)
	fmt.Println(result)
}
