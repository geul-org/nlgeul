//ff:func feature=token type=helper control=iteration dimension=1
//ff:what nodesOf — the nodes of a region, one for each span, in the order of the spans: the contest of a node holds the candidates of its span, best first by lambda, with the unit of the region's contest, its letters, and Glued as the region's for the node that begins the region and true for the others (at is where the region begins)
package token

import "slices"

func nodesOf(head Contest, at int, spans []Span, p pool) []Node {
	letters := []rune(head.Text)
	nodes := make([]Node, 0, len(spans))
	for _, sp := range spans {
		bs := slices.Clone(p.words[sp])
		sortByLambda(bs)
		c := Contest{Unit: head.Unit, Glued: sp.Start != at || head.Glued, Text: string(letters[sp.Start-at : sp.End-at]), Branches: bs}
		nodes = append(nodes, Node{Span: sp, Contest: c})
	}
	return nodes
}
