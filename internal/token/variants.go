//ff:func feature=token type=reader control=sequence
//ff:what variants — the ways a tail is read after the letters from b: as it is (unless its first piece was judged by a changed stem's own last syllable), once for each stem of the word table that the letters are a changed form of (when the tail begins with what that change happens before and the stem takes the tail), and once for each right form of the word table that the letters are a misspelt form of (misspelt: the right form takes the tail)
package token

// A variant for a stem of the table shows that stem (shown), so that
// the word is printed and looked up under it; a variant for a misspelt
// form shows the right form and keeps the letters (wrote). The order
// is: as it is, the changed forms, the misspelt forms.
func (sc *scan) variants(b int, t tail) []tail {
	var out []tail
	if t.irr == 0 {
		out = append(out, t)
	}
	if t.shown != "" || t.as != "" {
		return out
	}
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.head != 0 {
		out = sc.changedReads(out, stem, t)
	}
	if t.irr == 0 {
		out = sc.typoReads(out, stem, t)
	}
	return out
}
