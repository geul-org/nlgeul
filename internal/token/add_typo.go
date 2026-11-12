//ff:func feature=token type=reader control=sequence
//ff:what addTypo — put the entry of a closing written wrong into the index under its key with its graph stamped from the letters it is written with (and the rebuttal of a misspelt form); its count is not added to the total, nor its form to the closing set, and what it begins with is the right closing's
package token

// The count of the entry is the right closing's, which add has already
// put into the total once; the letters written wrong are not a closing
// the tables list, and no list of closings is told of them.
func (ix *index) addTypo(e *entry) {
	ix.edges += markerGraph(e, closedEnding(ix.words, e.form), ix.lexed)
	ix.graphs++
	k := entryKey(e)
	ix.byKey[k] = append(ix.byKey[k], e)
}
