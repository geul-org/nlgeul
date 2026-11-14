//ff:func feature=token type=helper control=iteration dimension=1
//ff:what count — how many ways of cutting go on from a span with the set of states a way has come to there: one when the span ends the portion, plus those of each span that follows and has a state after one of the set; counted once for each (span, set) (memo), exactly, as a big number
package token

import "math/big"

func (w *weave) count(node int, set []int, memo map[string]*big.Int) *big.Int {
	key := setKey(node, set)
	if v, ok := memo[key]; ok {
		return v
	}
	v := new(big.Int)
	if w.spans[node].End == w.to {
		v.SetInt64(1)
	}
	for _, nx := range w.succ[node] {
		if ns := w.next(set, nx); len(ns) > 0 {
			v.Add(v, w.count(nx, ns, memo))
		}
	}
	memo[key] = v
	return v
}
