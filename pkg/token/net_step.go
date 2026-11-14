//ff:func feature=token type=reader control=iteration dimension=1
//ff:what netStep — the states that end at place j (the places before it are worked out): each tail that ends at j grows what is kept before its stem (extend), then the signs between letters read as words of their own and the words of no markers, and the candidates are made states (settle); planted are states that stand at j before any of these (carried over from another scan: they keep their numbers of birth)
package token

func (sc *scan) netStep(best [][]*path, j int, planted []*path) {
	var cands []*path
	tails := sc.tailsAt(j)
	sc.marks[j] = markedFrom(tails)
	for _, t := range tails {
		cands = sc.extend(cands, best, t, j)
	}
	cands = sc.signWords(cands, best, j)
	best[j] = sc.settled(planted, sc.whole(cands, best, j, false))
}
