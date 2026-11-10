//ff:func feature=token type=reader control=sequence
//ff:what add — put one entry into the index under its key, with what it begins with noted and its graph stamped; its observed count is added to the total once for its form, however many readings the form has
package token

func (ix *index) add(e *entry) {
	e.head = headOf(e)
	ix.edges += markerGraph(e, closedEnding(ix.words, e.form), ix.lexed)
	ix.graphs++
	if ix.counted == nil {
		ix.counted = map[string]bool{}
	}
	if form := entryForm(e); !ix.counted[form] {
		ix.counted[form] = true
		ix.total += float64(e.count)
	}
	k := entryKey(e)
	ix.byKey[k] = append(ix.byKey[k], e)
}
