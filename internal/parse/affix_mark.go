//ff:func feature=parse type=reader control=sequence
//ff:what affixMark — a stem with "+" before 하, 되 or 시키 when a word stands before it (지시해 → 지시+해), or before an auxiliary verb's stem after a main verb's 어/아 form (해주 → 해+주)
package parse

// fused tells that a tail piece is fused into the stem's last syllable;
// the syllable is printed as it is, and only looked at without the
// piece.
func affixMark(stem []rune, fused bool) string {
	n := len(stem)
	base := stem[n-1]
	if fused {
		base = unfuse(base)
	}
	k := n
	if base == '하' || base == '되' {
		k = n - 1
	}
	if base == '키' && n >= 2 && stem[n-2] == '시' {
		k = n - 2
	}
	if k >= 1 && k < n {
		return leftMark(stem[:k]) + "+" + string(stem[k:])
	}
	if h := helperStart(stem, fused); h >= 1 {
		return aMark(stem[:h]) + "+" + string(stem[h:])
	}
	return string(stem)
}
