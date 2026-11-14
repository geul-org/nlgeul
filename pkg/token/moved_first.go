//ff:func feature=token type=helper control=sequence
//ff:what movedFirst — the first word of a candidate with its sign (the first part) taken off: no stem, so Attach is true; it keeps its verdict, its preference less the sign's own (signPref), and its end counted from the letters after the sign; false when no part is left
package token

func movedFirst(e step, head int, signPref float64) (step, bool) {
	w := e.word
	w.Parts = append([]Part{}, w.Parts[1:]...)
	if len(w.Parts) == 0 {
		return e, false
	}
	e.end -= head
	w.End, w.Attach = e.end, true
	e.pref -= signPref
	e.word, e.sig = w, w.String()
	return e, true
}
