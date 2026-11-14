//ff:func feature=token type=reader control=sequence
//ff:what settle — the candidates that end at a place put in order for the places after it: each last word is given its number of birth, then the candidates are made states of the net when the scan builds it (merged), else sorted best first with the same candidate once (top)
package token

func (sc *scan) settle(cands []*path) []*path {
	sc.numbered(cands)
	if sc.netted {
		return sc.merged(cands)
	}
	return sc.ix.top(cands)
}
