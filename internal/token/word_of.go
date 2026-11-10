//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wordOf — one word of a reading: the stem from b (a known word, an atom or a new word) and the markers after it as the parts of a word, with its verdict and preference
package token

import (
	"math"
	"slices"
	"strings"
)

// A stem the word table has stands as its kind's graph judges it; the
// force of "rarely seen" is set so that one observation still stands
// above a new word. A stem that a piece is fused into
// is printed, and looked up, with its last syllable as it was before
// fusing (마신 → 마시). In a stem an atom counts as one token for
// preference.
func (sc *scan) wordOf(b, end int, t tail) step {
	stem, name, v := sc.stemOf(b, t)
	e := step{end: end, word: Eojeol{Parts: []Part{{Text: stem, Stem: true}}, End: end}, verdict: sc.ix.unknown}
	if v > 0 {
		e.verdict, e.word.Parts[0].Names = v, namesOf(name)
	}
	e.pref = sc.ix.stemPref(stem, tailClass(t.pieces), v > 0, sc.tokens(b, t.start)) - sc.ix.structure(b, t.pieces)
	if t.as != "" {
		e.pref = t.weight - sc.ix.structure(b, t.pieces)
	}
	last := ""
	for _, p := range t.pieces {
		hidden, shown := hiddenOf(p.text)
		e.word.Parts = append(e.word.Parts, Part{Text: shown, Hidden: hidden, Names: namesOf(p.names), Unseen: p.unseen})
		e.verdict, e.pref, last = math.Min(e.verdict, p.verdict), e.pref+p.pref, p.names
	}
	e.sig = e.word.String()
	e.adn = strings.HasPrefix(last, "꾸밈:")
	e.conn = slices.Contains(connNames, last)
	if b > 0 && v > 0 && len(t.pieces) > 0 && slices.Contains(sc.ix.auxStems(), stem) {
		e.back = sc.ix.structure(b, nil) - sc.ix.structure(0, nil)
	}
	return e
}
