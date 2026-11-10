//ff:func feature=token type=reader control=iteration dimension=1
//ff:what variants — the ways a tail is read after the letters from b: as it is (unless its first piece was judged by a changed stem's own last syllable), and once for each stem of the word table that the letters are a changed form of, when the tail begins with what that change happens before and the stem takes the tail
package token

// A variant for a stem of the table shows that stem (shown), so that
// the word is printed and looked up under it.
func (sc *scan) variants(b int, t tail) []tail {
	var out []tail
	if t.irr == 0 {
		out = append(out, t)
	}
	if t.shown != "" || t.as != "" || t.head == 0 {
		return out
	}
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	for _, a := range sc.ix.alts[stem] {
		if _, v := sc.ix.known(a.stem, tailClass(t.pieces)); v > 0 && irrHeads[a.kind]&t.head != 0 && (t.irr == 0 || t.irr == a.kind) {
			read := t
			read.shown, read.irr = a.stem, 0
			out = append(out, read)
		}
	}
	return out
}
