//ff:func feature=token type=reader control=iteration dimension=1
//ff:what signBundles — the signs at one edge of a unit sent up, so that no letter of the sentence is dropped: each stretch of one sign in a row is a bundle of its own with the readings that stand for it (밥 먹었어? → ?[물음표], 글쎄... → ...[줄임표], #태그 → #[샵] and #[태그]), and stretches that no reading stands for are joined into one bundle of their letters with no name (->)
package token

// at is where the signs begin in the unit's letters, and after tells
// that they stand after the letters that are read.
func (sc *scan) signBundles(u unit, signs []rune, at int, after bool) []Bundle {
	var out []Bundle
	end := 0
	if after {
		end = len(sc.s)
	}
	for _, run := range sameRuns(signs) {
		text := string(signs[run[0]:run[1]])
		m := sc.edgeFacts(u, at+run[0], at+run[1], after, run[1] == len(signs))
		out = joinBare(out, sc.signBundle(text, m, end))
	}
	return out
}
