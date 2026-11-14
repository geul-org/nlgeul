//ff:func feature=token type=helper control=iteration dimension=1
//ff:what tiles — whether the nodes of a segment follow one another from the start of its region to its end, each node's letters the letters of its span (the node's contest has at least one candidate, each of one word), the letters of the nodes together the region's letters
package token

func tiles(sg Segment, letters []rune) bool {
	at := sg.Region.Start
	for _, n := range sg.Nodes {
		if n.Span.Start != at || string(letters[n.Span.Start:n.Span.End]) != n.Contest.Text || len(n.Contest.Branches) == 0 || !oneWordEach(n) {
			return false
		}
		at = n.Span.End
	}
	return at == sg.Region.End
}
