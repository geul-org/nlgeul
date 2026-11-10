//ff:func feature=token type=reader control=sequence
//ff:what add — put one entry into the index under its key, with its graph stamped
package token

func (ix *index) add(e *entry) {
	ix.edges += markerGraph(e, closedEnding(ix.words, e.form))
	ix.graphs++
	ix.total += float64(e.count)
	k := entryKey(e)
	ix.byKey[k] = append(ix.byKey[k], e)
}
