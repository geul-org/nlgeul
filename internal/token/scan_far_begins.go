//ff:func feature=token type=reader control=iteration dimension=1
//ff:what farBegins — the places before lo (farther from e than a new word's stem may be) from which a word that is not a new word may begin that ends at e: a key of the tables as it stands or with its last letter taken off, and the first word of a stem of words bound by a sign (boundBegins), the head of the unit (a sign of a pair or a span of code) whose word is not read as a new word; every place from which a candidate that is no new word can stand, and no more than that is needed: a place that is in them and stands for none only gives the new word, which is not made there
package token

import "slices"

func (sc *scan) farBegins(e, lo int) []int {
	var out []int
	if e == sc.head && e > 0 && lo > 0 {
		out = append(out, 0)
	}
	for b := max(0, e-sc.ix.keyMax); b < lo; b++ {
		if !sc.inside[b] && sc.farKey(b, e) {
			out = append(out, b)
		}
	}
	bound := sc.boundBegins(e, lo)
	if len(bound) == 0 {
		return out
	}
	out = append(out, bound...)
	slices.Sort(out)
	return slices.Compact(out)
}
