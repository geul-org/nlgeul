//ff:type feature=token type=model
//ff:what Segment — what a stream sends up each time: one way of cutting a stretch (Region) into words, the Rank-th best of the Ways ways there are (a big number: a stretch may have more ways than any integer holds), its words in order, and the sum of their best lambdas
package token

import "math/big"

// Ways is 1 for a stretch that is cut one way only (a word written
// apart); more than one is a contest of ways to cut it. Lambda is the sum
// over the nodes of the lambda of each node's best candidate.
type Segment struct {
	Region Span
	Rank   int
	Ways   *big.Int
	Nodes  []Node
	Lambda float64
}
