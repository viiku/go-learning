package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow, Extend the slice
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate new array
		// Grow by doubling,s for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a builtin function
	}
	z[len(x)] = y
	return z
}

// func main() {
// 	var runes []rune
// 	for _, r := range "Hello, World" {
// 		runes = append(runes, r)
// 	}
// 	fmt.Printf("%q\n", runes)
// }

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
