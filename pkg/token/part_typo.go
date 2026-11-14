//ff:func feature=token type=rule control=iteration dimension=1
//ff:what partTypo — the right form of a word of a bound stem that the word table gives as a wrong way to write another word, and how well the right form stands with what follows it (misspelt: the first right form, in the order of the table's lines, that stands; the word itself and 0 when none does)
package token

// Only the first right form is read, so that the candidates of a stem do
// not multiply with the right forms of its words.
func (ix *index) partTypo(part string, tail int) (string, float64) {
	for _, ty := range ix.typos[part] {
		if v := ix.misspelt(ty.right, ty.kinds, tail, false); v > 0 {
			return ty.right, v
		}
	}
	return part, 0
}
