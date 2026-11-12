//ff:func feature=token type=rule control=sequence
//ff:what partRight — one word of a bound stem as it is read with what follows it: the word itself and how well it stands (partVerdict), or, when it does not stand and typos is true, the right form of the first misspelt reading of it that stands (partTypo)
package token

func (ix *index) partRight(part string, tail int, typos bool) (string, float64) {
	if v := ix.partVerdict(part, tail); v > 0 || !typos {
		return part, v
	}
	return ix.partTypo(part, tail)
}
