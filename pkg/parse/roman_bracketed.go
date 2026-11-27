//ff:func feature=parse type=reader control=sequence
//ff:what romanBracketed — the candidate of a stem that is a Roman numeral written in the Latin capitals I, V and X, not left open by its letters (romanSure), that is the whole of a unit standing alone in a round bracket written right onto a word (bracketedAt: 크로뮴(III), 구리(II)): named 수:로마십진, with the verdict of a number
package parse

// A numeral in a round bracket right after a name tells which of the
// things of that name is meant: the oxidation number of an element is
// written so (the Stock nomenclature: iron(III)), with no space before
// the bracket.
func (sc *scan) romanBracketed(b int, stem string, t tail) []kindCandidate {
	if !sc.bracketed || b != 0 || t.start != len(sc.s) || len(t.pieces) > 0 || !romanSure(stem) {
		return nil
	}
	return []kindCandidate{{0, romanNumName, sc.ix.number}}
}
