// Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from
// the command line, invoking their preferred text editor when substantial text input is required.

package main

import (
	"log"
	"os"
	"time"
)

type User struct {
	Login   string
	HTMLURL string `json:"html_url"` // Fixed json tag
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	User      *User
	CreatedAt time.Time
	Title     string
	Body      string `json:"body,omitempty"`
	Assignees []string
	Milestone int
	Labels    []string
}

func getToken() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable not set")
	}
	return token
}

func createIssue() (*Issue, error) {
	Issue.Title = "test creating issue"
	return &Issue
}

func readIssue() {

}

func updateIssue() {

}

func deleteIssue() {

}

func main() {

}
