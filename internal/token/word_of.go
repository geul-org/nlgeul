//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wordOf — one word of a reading: the stem from b (a known word, an atom or a new word) and the markers after it, printed with "+" and the graph names, with its verdict and preference
package token

import (
	"math"
	"strings"
)

// A stem the word table has stands as its kind's graph judges it; the
// force of "rarely seen" is set so that one observation still stands
// above a new word. A stem that a piece is fused into
// is printed, and looked up, with its last syllable as it was before
// fusing (마신 → 마시). In a stem an atom counts as one token for
// preference.
func (sc *scan) wordOf(b, end int, t tail) eojeol {
	stem, name, v := sc.stemOf(b, t)
	e := eojeol{end: end, marked: stem, names: newWord, verdict: sc.ix.unknown}
	if v > 0 {
		e.names, e.verdict, e.marked = name, v, stem+"["+name+"]"
	}
	e.pref = sc.ix.stemPref(stem, v > 0, sc.tokens(b, t.start)) - sc.ix.structure(b, t.pieces)
	if t.as != "" {
		e.pref = t.weight - sc.ix.structure(b, t.pieces)
	}
	names := []string{e.names}
	for _, p := range t.pieces {
		e.marked += "+" + p.text + "[" + p.names + "]"
		e.verdict, e.pref = math.Min(e.verdict, p.verdict), e.pref+p.pref
		names = append(names, p.names)
	}
	e.names = strings.Join(names, " + ")
	e.adn = len(t.pieces) > 0 && strings.HasPrefix(names[len(names)-1], "꾸밈:")
	return e
}
