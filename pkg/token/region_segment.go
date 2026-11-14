//ff:func feature=token type=helper control=iteration dimension=1
//ff:what segment — the way made of the nodes given by their places as what is sent up: the nodes in order, their best lambdas summed, and the rank (the next of those made so far)
package token

func (rg *region) segment(ids []int) Segment {
	seg := Segment{Region: rg.span, Rank: len(rg.out) + 1, Ways: rg.ways}
	for _, id := range ids {
		seg.Nodes = append(seg.Nodes, rg.nodes[id])
		seg.Lambda += rg.best[id]
	}
	return seg
}
