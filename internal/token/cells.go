//ff:type feature=token type=model
//ff:what cells — a stretch of a unit read by the net, ready to become a region: where the stretch begins in the letters of the unit, its cells (word-places) in the order of their spans, the pairs of words that stand side by side, how many ways of cutting there are (a big number), whether a line of spans is one of them, the places where a word printed the same came with other values (clashes), and how many (word-place, set of states) keys the counting of ways made (sets)
package token

import "math/big"

type cells struct {
	from    int
	nodes   []cell
	edges   []edge
	ways    *big.Int
	member  func([]Span) bool
	clashes int
	sets    int
}
