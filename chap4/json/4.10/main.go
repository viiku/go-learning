// Exercise 4.10: Modify issues to report the results in age categories, say less than a month old,
// less than a year old, and more than a year old.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"` // Fixed json tag
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// SearchIssues queries the GitHub issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	// fmt.Print(resp)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()

	var lessThanMonth, lessThanYear, moreThanYear []*Issue

	for _, item := range result.Items {
		age := now.Sub(item.CreatedAt)

		switch {
		case age < 30*24*time.Hour:
			lessThanMonth = append(lessThanMonth, item)
		case age < 365*24*time.Hour:
			lessThanYear = append(lessThanYear, item)
		default:
			moreThanYear = append(moreThanYear, item)
		}
	}

	// Display issues in each category
	fmt.Println("\nIssues less than a month old:")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nIssues less than a Year old:")
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nIssues more than a Year old:")
	for _, item := range moreThanYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
