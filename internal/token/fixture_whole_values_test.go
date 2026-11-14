//ff:func feature=token type=helper control=iteration dimension=1
//ff:what wholeValues — the candidates (branchValues) of the node of a region that is the whole region, and whether there is one
package token

func wholeValues(rg *region) ([]string, bool) {
	for _, n := range rg.nodes {
		if n.Span == rg.span {
			return branchValues(n.Contest), true
		}
	}
	return nil, false
}
