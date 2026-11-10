//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addShort — the index entries of one shortened closing: named with the form it is short for, read only after a predicate's stem of the word table and not right before a particle, with its kind of tail; the copula's graph is told when it does not take it
package token

func (ix *index) addShort(row short) {
	for _, e := range closingEntries(row.m) {
		e.name += "(" + row.origin + ")"
		e.wide, e.gate, e.tail, e.bars = true, ix.pred, row.tail, row.bars
		ix.add(e)
		if e.bars {
			ix.barred = append(ix.barred, e.name)
		}
	}
}
