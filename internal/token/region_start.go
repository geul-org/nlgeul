//ff:func feature=token type=helper control=iteration dimension=1
//ff:what start — put the words that begin the region in the frontier, one walk each, in the order of the nodes
package token

func (rg *region) start() {
	for i, n := range rg.nodes {
		if n.Span.Start == rg.span.Start {
			rg.front.push(&walk{node: i, f: rg.rest[i]})
		}
	}
}
