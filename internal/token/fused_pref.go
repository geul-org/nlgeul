//ff:func feature=token type=reader control=sequence
//ff:what fusedPref — the preference of a fused 어/아 or ㅆ as the whole form it is written out as (가 → 가+아: that of 아; 갔 → 가+았: that of 았), when the index has that form with a count; else the preference given
package token

func (ix *index) fusedPref(e *entry, piece string, pref float64) float64 {
	if e.lead != leadVowel && e.lead != jongSS {
		return pref
	}
	whole := []rune(piece + e.form)
	w := ix.named(whole[len(whole)-1], labelOf(e.name)+":"+string(whole))
	if w.count == 0 {
		return pref
	}
	return ix.pref(w.count, string(whole))
}
