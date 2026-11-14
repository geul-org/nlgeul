//ff:func feature=token type=helper control=iteration dimension=2
//ff:what nodesAt — the nodes at a span, from all the ways of a stretch
package token

func nodesAt(segs []Segment, span Span) []Node {
	var out []Node
	for _, sg := range segs {
		for _, n := range sg.Nodes {
			if n.Span == span {
				out = append(out, n)
			}
		}
	}
	return out
}
