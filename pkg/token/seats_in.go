//ff:func feature=token type=helper control=iteration dimension=1
//ff:what seatsIn — the nodes of one region's best way as seats of the k-th region, each with how its first candidate reads it
package token

func seatsIn(k int, nodes []Node) []seat {
	out := make([]seat, 0, len(nodes))
	for _, n := range nodes {
		out = append(out, seat{region: k, node: n, near: nearOf(n.Contest.Branches[0])})
	}
	return out
}
