package main

import "golang.org/x/net/html"

// dfsVisit differs from visit in that it traverses the n.FirstChild
// linked list using recursive calls instead of a loop
func dfsVisit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	// Just for testing pre-order, in-order, and post-order traversal
	//links = dfsVisit(links, n.FirstChild)
	//links = dfsVisit(links, n.NextSibling)

	if  n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = dfsVisit(links, n.FirstChild)
	links = dfsVisit(links, n.NextSibling)

	return links
}
