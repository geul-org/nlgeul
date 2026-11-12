//ff:func feature=token type=reader control=iteration dimension=1
//ff:what grown — the readings kept at b, each grown by the word from b to j that tail t makes, once for each way its stem stands (a kind, a set, a name; growWord), when the word before the tail passes the tail's graph for that kind, the stem's last syllable is one it may have (restored), and the stem stands (a new word as its graph judges it)
package token

func (sc *scan) grown(cands []*path, kept []*path, t tail, b, j int) []*path {
	if !sc.restored(b, t) {
		return cands
	}
	for _, e := range sc.stepsOf(b, j, t) {
		if sc.gated(b, t, e) {
			cands = sc.growWord(cands, kept, t, e, b, j)
		}
	}
	return cands
}
