//ff:func feature=token type=reader control=sequence
//ff:what TestSameSpans — nodes stand at exactly the spans given: the same, one span fewer, another span
package token

import "testing"

func TestSameSpans(t *testing.T) {
	nodes := []Node{{Span: Span{0, 2}}, {Span: Span{2, 5}}}
	if !sameSpans(nodes, []Span{{0, 2}, {2, 5}}) || sameSpans(nodes, []Span{{0, 2}}) || sameSpans(nodes, []Span{{0, 2}, {2, 4}}) {
		t.Fatal("spans")
	}
}
