package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var a, b int
	_, err := fmt.Scanf("%d%d", a, b)
	if err != nil {
		// var error = fmt.Errorf("Please enter both values")
		return fmt.Printf("The Error %s\n", err)
	}
	var res int = gcd(a, b)
	print(res)
}
