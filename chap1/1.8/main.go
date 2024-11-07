// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix.

package main

import (
	"fmt"
	"net/http"
	"os"
)

func HasPrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func main() {
	for _, url := range os.Args[1:] {
		if !HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Printf("%s", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", resp.Body)
		// b, err := io.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// }
		// fmt.Printf("%s", b)
	}
}
