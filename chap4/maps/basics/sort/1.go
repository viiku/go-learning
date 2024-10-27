package main

import (
	"fmt"
)

func main() {
	// a map literal to create a new map populated
	// with some initial key/value pairs:
	ages := map[string]int{
		"charlie": 34,
		"apsks":   516,
		"alice":   31,
	}
	fmt.Print(ages, "\n")
	fmt.Printf("%d", len((ages)))
	// // creating empty array of strings
	// var names []string
	// fmt.Printf("%s\n", names)
	// for name := range ages {
	// 	names = append(names, name)
	// }
	// sort.Strings(names)
	// for _, name := range names {
	// 	fmt.Printf("%s\t%d\n", name, ages[name])
	// }
	names := make([]string, 0, len(ages))
	fmt.Printf("%d\n", len(names))
}
