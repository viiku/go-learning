// squares returns a function that returns
// the next square number each time it is called.

package main

import "fmt"

func Squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := Squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
