//ff:func feature=token type=helper control=iteration dimension=1
//ff:what ways — how many ways of cutting the portion there are: the lines of spans from where it begins to where it ends that chain (see member), counted exactly as a big number over (span, set of states) keys (every state at a span that begins the portion is the first set)
package token

import "math/big"

func (w *weave) ways() *big.Int {
	total := new(big.Int)
	memo := map[string]*big.Int{}
	for node, sp := range w.spans {
		if sp.Start == w.from {
			total.Add(total, w.count(node, w.of[node], memo))
		}
	}
	w.sets = len(memo)
	return total
}
