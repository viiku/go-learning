package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Step 1
	// inputFileData, err := os.ReadFile("name.txt")
	// if err != nil {
	// 	fmt.Println("Error Reading file")
	// 	return
	// }

	// Step 2
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Error Reading file")
		return
	}

	defer file.Close()

	wordFreq := make(map[string]int)

	// Set up a scanner to read the file word by word
	scanner := bufio.NewScanner(file)
	// Set scanner to split by words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)        // Convert to lowercase to avoid case sensitivity
		word = strings.Trim(word, ".,!?;:") // Trim punctuation
		wordFreq[word]++                    // Increment word frequency
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for word, count := range wordFreq {
		fmt.Printf("%s %d\n", word, count)
	}
}
