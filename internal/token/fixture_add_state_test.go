//ff:func feature=token type=helper control=iteration dimension=1
//ff:what addState — a state made up for the net tests: the word given grown from each state before it (each a candidate), merged into the one state of the net that stands for them; the first state before it gives where the word begins
package token

func addState(sc *scan, before []*path, e step) *path {
	var cands []*path
	for _, p := range before {
		cands = append(cands, grow(p, e))
	}
	return sc.merged(cands)[0]
}
