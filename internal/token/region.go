//ff:type feature=token type=model
//ff:what region — a stretch of a sentence and the ways of cutting it into words, kept while a stream hands the ways out one at a time: its span, its nodes (a word for each span any way puts one at), the best lambda of each node and the most a way can reach from it, the nodes that follow each, the ways of cutting there are (whether a line of spans is one of them: member, and how many there are: ways, a big number), the walks not yet taken further, the segments made so far and how many of them were sent, the ways the search made by joining nodes that no cut has (skipped), and the places where a word printed the same came with other values (clashes)
package token

import "math/big"

type region struct {
	span    Span
	nodes   []Node
	best    []float64
	rest    []float64
	next    [][]int
	member  func([]Span) bool
	ways    *big.Int
	front   frontier
	out     []Segment
	sent    int
	skipped int
	clashes int
}
