package main

import (
	"fmt"
)

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *Node) []string {
	if n.Type == ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = visit(links, c)
		}
	}
	return links
}

func linkFinder(url string) {
	urlData, err := fetchURLData(url)
	if err != nil {
		fmt.Print("Not able to fetch URL Data.")
	}
	doc, err := Parse(urlData)
	if err != nil {
		fmt.Print("Not able to Parse URL Data.")
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	// return link
}
