//ff:func feature=token type=reader control=selection
//ff:what copulaOf — the copula candidate a predicate tail may have: where the noun before it ends, and whether the stem 이 is said (a syllable 이, or the syllable a piece is fused into) or may be dropped; whether a vowel stands before a dropped stem and what it is dropped before (a closing right after it, not a slot) are left to the dropped stem's graph; a tail that asks for a predicate's stem or for a changed stem has no such candidate
package token

// The third value is false when the tail has no such candidate.
func (sc *scan) copulaOf(t tail) (int, *entry, bool) {
	whole := t.base == 0
	switch {
	case t.close == nil || t.close.holds || t.irr != 0:
		return 0, nil, false
	case !whole && t.base == '이':
		return t.start - 1, sc.ix.copula, true
	case whole && t.start >= 1 && sc.s[t.start-1] == '이':
		return t.start - 1, sc.ix.copula, true
	case whole && t.start >= 1:
		return t.start, sc.ix.dropped, true
	case !whole && t.gate == nil:
		return t.start, sc.ix.dropped, true
	}
	return 0, nil, false
}
