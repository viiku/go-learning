package main

import "fmt"

func main() {
	var pc [256]byte
	for _, v := range pc {
		fmt.Printf("%d", v)
	}
}
