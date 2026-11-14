//ff:func feature=token type=reader control=sequence
//ff:what keptBefore — the candidates a word from b up to end grows: those that end at b, and when the word is the one letters of the unit (no fused piece, not shown) that the cues show as a word that follows a word no table lists (shows), the candidates kept at b where a sign stands (those that end at b and those that end with a word of no markers there, keptAt)
package token

func (sc *scan) keptBefore(best [][]*path, b, end int, t tail) []*path {
	if t.base == 0 && t.shown == "" && t.as == "" && sc.shows[b][end] {
		return sc.keptAt(best, b)
	}
	return best[b]
}
