//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sameSpans — whether the nodes stand at exactly the spans given, in order
package token

func sameSpans(nodes []Node, want []Span) bool {
	if len(nodes) != len(want) {
		return false
	}
	for i, nd := range nodes {
		if nd.Span != want[i] {
			return false
		}
	}
	return true
}
