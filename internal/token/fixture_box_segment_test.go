//ff:func feature=token type=helper control=sequence
//ff:what boxSegment — a segment made up for the box tests: its rank, how many ways its stretch has, its lambda, its place and its nodes
package token

import "math/big"

func boxSegment(rank, ways int, lambda float64, region Span, nodes ...Node) Segment {
	return Segment{Region: region, Rank: rank, Ways: big.NewInt(int64(ways)), Nodes: nodes, Lambda: lambda}
}
