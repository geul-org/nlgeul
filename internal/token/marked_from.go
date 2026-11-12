//ff:func feature=token type=reader control=iteration dimension=1
//ff:what markedFrom — the places from which the letters up to where some tails end read as markers of the index alone: where a tail of particles, slots and closings starts, when nothing of it is fused into the stem before and it names no stem of its own, nor a marker written wrong (the letters of one are a word of their own as well)
package token

func markedFrom(tails []tail) map[int]bool {
	out := map[int]bool{}
	for _, t := range tails {
		if t.base == 0 && t.only == "" && t.as == "" && len(t.pieces) > 0 && !t.pieces[0].unseen && t.pieces[0].tail != 0 && !anyWrote(t.pieces) {
			out[t.start] = true
		}
	}
	return out
}
