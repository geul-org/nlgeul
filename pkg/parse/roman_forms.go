//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what romanForms — the candidates of a stem that is one Roman numeral letter of Unicode (the shape 로마 숫자 of the tables of special letters: Ⅲ, ⅳ) standing alone, with nothing written onto it, read by the letter's shape: one for each name the tables give the letter (Ⅲ[로마 숫자:수]), with the verdict of a number, put before the reading of the letter as a number (Ⅲ[수:로마십진]); letters in a row (ⅩⅢ), a letter with a unit or a particle on it (Ⅷ세, Ⅲ의) and a numeral written in Latin capitals (XII) are the number alone
package parse

func (ix *index) romanForms(stem string, t tail) []kindCandidate {
	r := []rune(stem)
	if len(r) != 1 || len(t.pieces) > 0 || !formShaped(r[0], romanShape) {
		return nil
	}
	var out []kindCandidate
	for _, name := range formNames(r[0]) {
		out = append(out, kindCandidate{0, name, ix.number})
	}
	return out
}
