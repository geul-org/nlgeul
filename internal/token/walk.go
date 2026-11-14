//ff:type feature=token type=model
//ff:what walk — a way through the words of a region taken so far, as a chain: the last word (its place in the region's nodes), the walk before it, the lambdas of the words before the last summed, and the most the whole way can reach (f)
package token

type walk struct {
	prev *walk
	node int
	got  float64
	f    float64
}
