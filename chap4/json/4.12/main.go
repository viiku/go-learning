// Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
// many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
// that, using this index, prints the URL and transcript of each comic that matches a search term
// provided on the command line.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const baseURL = "https://xkcd.com"
const indexFile = "info.0.json"

type Comic struct {
	Num         int    `json:"num"`
	Title       string `json:"title"`
	Safe_Title  string `json:"safe_title"`
	Month       string `json:"month"`
	Year        string `json:"year"`
	URL         string `json:"url"`
	Description string `json:"transcript"`
}

func fetchComicData(num int) (*Comic, error) {

	url := fmt.Sprintf("%s/%d/info.0.json", baseURL, num)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	comic.URL = fmt.Sprintf("%s/%d", baseURL, comic.Num)
	return &comic, nil
}

func buildIndex() error {
	var comics []Comic

	// Fetch a range of comics e.g., from 1 to 100
	for i := 1; i <= 100; i++ {
		comic, err := fetchComicData(i)
		if err != nil {
			log.Printf("Error fetching comic %d: %v\n", i, err)
			continue
		}
		comics = append(comics, *comic)

		// Save index to file
		data, err := json.Marshal(comics)
		if err != nil {
			return fmt.Errorf("failed to encode index: %w", err)
		}
		if err := os.WriteFile(indexFile, data, 0644); err != nil {
			return fmt.Errorf("failed to write index to file: %w", err)
		}
	}
	fmt.Println("Index built and saved successfully.")
	return nil
}

func loadIndex() ([]Comic, error) {
	data, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	var comics []Comic
	if err := json.Unmarshal(data, &comics); err != nil {
		return nil, fmt.Errorf("failed to decode index JSON: %w", err)
	}
	return comics, nil
}

func searchComics(comics []Comic, term string) {
	term = strings.ToLower(term)
	found := false

	for _, comic := range comics {
		if strings.Contains(strings.ToLower(comic.Description), term) ||
			strings.Contains(strings.ToLower(comic.Title), term) {
			fmt.Printf("Title: %s\nURL: %s\nTranscript: %s\n\n", comic.Title, comic.URL, comic.Description)
			found = true
		}
	}

	if !found {
		fmt.Println("No comics found with the given term.")
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <search term>")
	}

	term := strings.Join(os.Args[1:], " ")

	// If index file doesn't exist, build it
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		fmt.Println("Index file not found. Building index...")
		if err := buildIndex(); err != nil {
			log.Fatalf("Failed to build index: %v", err)
		}
	}
	// Load index and search
	comics, err := loadIndex()
	if err != nil {
		log.Fatalf("Failed to load index: %v", err)
	}

	searchComics(comics, term)
}
