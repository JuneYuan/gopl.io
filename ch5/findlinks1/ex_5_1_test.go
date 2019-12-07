package main

import (
	"testing"

	"golang.org/x/net/html"
)

func Test_dfsVisit(t *testing.T) {
	type args struct {
		links []string
		n     *html.Node
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case - 1",
			args: args{
				links: nil,
				n:     mockHtmlTree(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := visit(tt.args.links, tt.args.n)
			t.Logf("dfsVisit() = %v, want %v", got, tt.want)
			//if got := dfsVisit(tt.args.links, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("dfsVisit() = %v, want %v", got, tt.want)
			//}
		})
	}
}

// mock an html tree whose:
// pre-order is [A B E C F G D]
// in-order is [E B F G C D A]
// post-order is [E G F D C B A]
func mockHtmlTree() *html.Node {
	A := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "A"}}}
	B := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "B"}}}
	C := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "C"}}}
	D := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "D"}}}
	E := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "E"}}}
	F := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "F"}}}
	G := &html.Node{Type: html.ElementNode, Data: "a", Attr: []html.Attribute{{Key: "href", Val: "G"}}}

	A.FirstChild = B
	B.NextSibling = C
	C.NextSibling = D
	B.FirstChild = E
	C.FirstChild = F
	F.NextSibling = G

	return A
}
