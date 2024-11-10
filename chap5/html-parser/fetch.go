package main

import (
	"fmt"
	"io"
	"net/http"
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
