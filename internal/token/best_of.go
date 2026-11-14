//ff:func feature=token type=helper control=iteration dimension=1
//ff:what bestOf — the lambda of the best candidate of each node (the candidates of a node are kept best first)
package token

func bestOf(nodes []Node) []float64 {
	out := make([]float64, 0, len(nodes))
	for _, n := range nodes {
		out = append(out, lambdaOf(n.Contest.Branches[0]))
	}
	return out
}
