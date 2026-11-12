//ff:func feature=token type=reader control=sequence
//ff:what onStem — the particle readings that end at j as they stand right on a stem, by form (as byForm gives them): those given, unless the unit ends at j and a sign that ends a sentence closes it; there the graphs are asked once more, told so, and the readings that still stand come back (문제야. keeps the vocative 야 and not the auxiliary one)
package token

// A particle after another particle or after a closing is not asked
// so: 에서야. and 하고 나서야. keep their auxiliary 야.
func (sc *scan) onStem(j, next int, merged map[string]piece, kinds map[string][]*entry) (map[string]piece, map[string][]*entry) {
	if !sc.closed || j != len(sc.s) {
		return merged, kinds
	}
	pieces, from, _ := sc.groupAt(classParticle, j, next, false, true)
	_, merged, kinds = byForm(pieces, from)
	return merged, kinds
}
