package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchURLData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("fetch: %v", err)
	}

	// Read all of the response body
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("fetch: reading %s: %v", url, err)
	}
	respBody := string(b[:])
	return respBody, nil
}

func linkFinder(inputURL string) ([]string, error) {
	urlData, err := fetchURLData(inputURL)
	if err != nil {
		fmt.Print(err)
	}

	doc, err := html.Parse(io.Reader)
	if err != nil {
		fmt.Print("Not able to Parse URL Data.")
	}

}

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(url)
		linkFinder(url)
	}
}
