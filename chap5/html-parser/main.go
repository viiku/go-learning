package main

import (
	"fmt"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(url)
		linkFinder(url)
	}
}
