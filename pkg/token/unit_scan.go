//ff:func feature=token type=reader control=sequence
//ff:what unitScan — the scan of one unit before it is read: the letters that are read (coreOf), the unit's head, and whether a sign that ends a sentence stands right after those letters in the unit (closes)
package token

func unitScan(ix *index, u unit) *scan {
	sc := &scan{ix: ix, s: coreOf(u), head: u.head, headName: u.name}
	sc.closed = closes(sc.signsAfter(u.text))
	return sc
}
