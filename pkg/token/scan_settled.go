//ff:func feature=token type=reader control=sequence
//ff:what settled — the states that end at a place: the states planted there (carried over, with the numbers of birth they have) and then the candidates the scan made, each last word given its number of birth and the candidates made states; without planted ones, settle
package token

func (sc *scan) settled(planted, cands []*path) []*path {
	if len(planted) == 0 {
		return sc.settle(cands)
	}
	sc.numbered(cands)
	return sc.merged(append(append([]*path{}, planted...), cands...))
}
