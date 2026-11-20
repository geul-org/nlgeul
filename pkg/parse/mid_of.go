//ff:func feature=parse type=helper control=sequence
//ff:what midOf — the sign between letters that a word stands for: a word of one part named midName only, at the place where the run of the sign that ends at its end begins (midStart); false for any other word
package parse

func (sc *scan) midOf(w Eojeol) (midSign, bool) {
	if len(w.Parts) != 1 || w.Parts[0].Name.Kind != midName {
		return midSign{}, false
	}
	m, ok := sc.mid[sc.midStart(w.End)]
	return m, ok
}
