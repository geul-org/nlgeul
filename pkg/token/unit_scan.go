//ff:func feature=token type=reader control=sequence
//ff:what unitScan — the scan of one unit before it is read: the letters that are read (coreOf), the unit's head, whether a sign that ends a sentence stands right after those letters in the unit (closes), and whether those letters are the sound of Han characters glossed in a round bracket right after them (glossedBy)
package token

func unitScan(ix *index, u unit) *scan {
	sc := &scan{ix: ix, s: coreOf(u), head: u.head, headName: u.name}
	after := sc.signsAfter(u.text)
	sc.closed, sc.glossed = closes(after), glossedBy(sc.s[sc.head:], u.gloss, after)
	return sc
}
