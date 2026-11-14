//ff:func feature=token type=helper control=iteration dimension=2
//ff:what statesAfter — the states at a span that have a state of the set before them (all of them for the first span), in the order given
package token

import "slices"

func statesAfter(states []*path, sp Span, set []*path, first bool) []*path {
	var out []*path
	for _, p := range states {
		if p.start != sp.Start || p.last.end != sp.End {
			continue
		}
		for _, b := range p.preds {
			if first || slices.Contains(set, b) {
				out = append(out, p)
				break
			}
		}
	}
	return out
}
