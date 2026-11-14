//ff:func feature=token type=reader control=sequence
//ff:what entryPiece — the piece one entry read at one place makes, with the verdict its graph gave: the letters as they stand, the entry's name, its kind of tail and its preference; for an entry that reads a closing written wrong the right closing as the text (with its letters kept in wrote) and the preference of the right closing's form
package token

// The count of an entry that reads a closing written wrong is the right
// closing's, so its preference is that of the right closing's form (what
// stands in brackets in right is not a letter of the form).
func (sc *scan) entryPiece(e *entry, text string, v float64) piece {
	form := e.form
	if e.wrote != "" {
		_, form = hiddenOf(e.right)
		text = e.right
	}
	return piece{text, e.name, v, sc.ix.pref(e.count, form), tailOf(e), false, e.wrote}
}
