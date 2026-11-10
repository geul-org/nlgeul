//ff:func feature=token type=reader control=iteration dimension=1
//ff:what addAlt — note one changed form of one stem; for ㄷ and ㅅ also the stem's own last syllable under the changed one, once
package token

func (ix *index) addAlt(form, stem string, kind int) {
	ix.alts[form] = append(ix.alts[form], alt{stem, kind})
	if kind != irrD && kind != irrS {
		return
	}
	f, s := []rune(form), []rune(stem)
	last, own := f[len(f)-1], string(s[len(s)-1])
	for _, a := range ix.irrEnds[last] {
		if a.stem == own && a.kind == kind {
			return
		}
	}
	ix.irrEnds[last] = append(ix.irrEnds[last], alt{own, kind})
}
