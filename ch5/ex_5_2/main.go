package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var (
	f = count()
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise 5.2: %v\n", err)
		os.Exit(1)
	}

	m := make(map[string]int)
	populate(m, doc)
	display(m)
}

func populate(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		if _, ok := m[n.Data]; ok {
			m[n.Data] += 1
		} else {
			m[n.Data] = 1
		}
		// just debugging:
		// fmt.Printf("#%d: Type=%v, Data=%v, Attr=%v\n", f(), n.Type, n.Data, n.Attr)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		populate(m, c)
	}
}

func display(m map[string]int) {
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

// count is for debugging
func count() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}