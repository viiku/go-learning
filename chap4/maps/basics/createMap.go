package main

import "fmt"

func main() {
	// Empty map
	ages := map[string]int{}
	fmt.Println(ages)
	// adding some values
	ages["alice"] = 31
	ages["ppk"] = 32
	fmt.Println(ages)
	// printing values
	ages["alice"] = 32
	fmt.Println(ages["alice"])
	// deleting keys
	delete(ages, "alice")
	// returns zero if not present
	fmt.Println(ages["alice"])
	ages["alice"] = 35
	ages["sksk"] = 33
	ages["slsl"] = 34
	for name, v := range ages {
		fmt.Printf("%s\t%d\n", name, v)
	}
}
