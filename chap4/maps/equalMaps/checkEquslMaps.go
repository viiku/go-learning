package main

import "fmt"

func isEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		fmt.Printf("%s %d\n", k, xv)
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	ages1 := make(map[string]int)
	ages2 := make(map[string]int)
	ages1["alice"] = 31
	ages1["bob"] = 32
	ages1["pk"] = 33
	ages1["test"] = 34

	ages2["alice"] = 31
	ages2["bob"] = 32
	ages2["pk"] = 34
	ages1["test"] = 35

	sss := isEqual(ages1, ages2)
	fmt.Println(sss)

}
