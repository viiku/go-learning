// Without initialized variables gets 0 and empty values respectivly
package main

import (
	"fmt"
)

const pi = 3.14

func callingFunction() {
	// Short Variable Declaration
	// used mostly inside functions
	freq := pi * 3.0
	fmt.Printf("%f\n", freq)
}
func main() {
	var name int
	var game string
	fmt.Printf("%d\n%s\n", name, game)
	var i, j, k int
	fmt.Printf("%d%d%d\n", i, j, k)
	var integer, floating, stringssss = 1, 2.53, "vivek"
	fmt.Printf("%d%f%s\n", integer, floating, stringssss)

	callingFunction()
}
