// Implement unix basename command in go
package main

import (
	"fmt"
)

func basename(s string) string {
	// basename removes directory components and a .suffix.
	// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	var str, name string
	fmt.Scan(&str)
	name = basename(str)
	fmt.Println(name)
}
