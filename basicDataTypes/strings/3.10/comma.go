// insert commas at every third place from right to left in a number:
package main

import (
	"bytes"
	"fmt"
)

func insertComma(str string) string {
	var buf bytes.Buffer

	// Get length of string
	n := len(str)

	// Calculate number of digits remaining before first comma
	remainingDigits := n % 3
	if remainingDigits == 0 {
		remainingDigits = 3
	}

	// Process each character
	for i, v := range str {
		if i == remainingDigits {
			buf.WriteByte(',')
			remainingDigits += 3
		}
		buf.WriteRune(v)
	}

	return buf.String()
}

func main() {
	var str string
	fmt.Print("Enter a number: ")
	fmt.Scan(&str)

	result := insertComma(str)
	fmt.Println("Result:", result)
}
