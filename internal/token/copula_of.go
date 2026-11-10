//ff:func feature=token type=reader control=selection
//ff:what copulaOf — the copula reading a predicate tail may have: where the noun before it ends, and whether the stem 이 is said (a syllable 이, or the syllable a piece is fused into) or dropped after a vowel before one of the closings that allow it; a tail that asks for a predicate's stem or for a changed stem has no such reading
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The third value is false when the tail has no such reading.
func (sc *scan) copulaOf(t tail) (int, *entry, bool) {
	whole := t.base == 0
	switch {
	case t.close == nil || t.close.holds || t.irr != 0:
		return 0, nil, false
	case !whole && t.base == '이':
		return t.start - 1, sc.ix.copula, true
	case whole && t.start >= 1 && sc.s[t.start-1] == '이':
		return t.start - 1, sc.ix.copula, true
	case (len(t.pieces) != 1 && t.pieces[0].names != jiName) || !(slices.Contains(droppable, t.close.name) || (sc.ix.lexed && slices.Contains(droppableKnown, t.close.name))):
		return 0, nil, false
	case whole && t.start >= 1 && rule.JongOf(sc.s[t.start-1]) == 0:
		return t.start, sc.ix.dropped, true
	case !whole && t.gate == nil:
		return t.start, sc.ix.dropped, true
	}
	return 0, nil, false
}
