// Insert comma at every third place
// 12345 -> 12,345
package main

import "fmt"

func comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return comma(str[:n-3]) + "," + str[n-3:]
}

func main() {
	var number, result string
	fmt.Scan(&number)
	result = comma(number)
	fmt.Println(result)
}
