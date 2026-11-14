//ff:func feature=token type=helper control=sequence
//ff:what midOf — the sign between letters that a word stands for: a word of one part named midName only, at the place its end less one gives; false for any other word
package token

func (sc *scan) midOf(w Eojeol) (midSign, bool) {
	if len(w.Parts) != 1 || w.Parts[0].Name.Kind != midName {
		return midSign{}, false
	}
	m, ok := sc.mid[w.End-1]
	return m, ok
}
