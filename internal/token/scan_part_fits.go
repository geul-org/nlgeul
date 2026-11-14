//ff:func feature=token type=reader control=sequence
//ff:what partFits — whether the letters from b to e may be one word of a bound stem that stands: a key of the tables, or a run of Latin letters and digits (more than partVerdict and partTypo let stand, so that nothing is left out)
package token

func (sc *scan) partFits(b, e int) bool {
	text := sc.letters(b, e)
	return sc.ix.keyed(text) || runOnly(text)
}
