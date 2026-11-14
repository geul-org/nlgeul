//ff:func feature=token type=reader control=sequence
//ff:what lastPartFits — whether the letters from b to e may be the last word of a bound stem that stands: no longer than the longest key (its last letter may stand for the one a piece was fused into), or a run of Latin letters and digits
package token

func (sc *scan) lastPartFits(b, e int) bool {
	return e-b <= sc.ix.keyMax || runOnly(sc.letters(b, e))
}
