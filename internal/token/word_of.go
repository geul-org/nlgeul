//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wordOf — one word of a reading: the stem from b (a known word, a known word written wrong, an atom or a new word) and the markers after it as the parts of a word, with its verdict and preference; a stem no table lists stands as the new word's graph judges it, and a word whose stem does not stand has a verdict of 0 or less
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
// preference. The boundary before an auxiliary verb's stem is given
// back only when the stem stands as a predicate (주 before the modifying
// 는, not 주 the noun before the particle 는).
func (sc *scan) wordOf(b, end int, t tail) step {
	stem, name, v, wrote := sc.stemOf(b, t, b == 0 && end == len(sc.s))
	e := step{end: end, word: Eojeol{Parts: []Part{{Text: stem, Stem: true}}, End: end}, verdict: v, typo: wrote != ""}
	if v > 0 {
		e.word.Parts[0].Names, e.word.Parts[0].Written = namesOf(name), wrote
	} else {
		e.verdict = sc.unlisted(b, end, stem, t)
	}
	e.pref = sc.ix.stemPref(stem, tailClass(t.pieces), v > 0, sc.tokens(b, t.start)) - sc.ix.structure(b, t.pieces)
	if t.as != "" {
		e.pref = t.weight - sc.ix.structure(b, t.pieces)
	}
	last := ""
	for _, p := range t.pieces {
		hidden, shown := hiddenOf(p.text)
		e.word.Parts = append(e.word.Parts, Part{Text: shown, Hidden: hidden, Names: namesOf(p.names), Unseen: p.unseen, Written: p.wrote})
		e.verdict, e.pref, last = math.Min(e.verdict, p.verdict), e.pref+p.pref, p.names
		e.typo = e.typo || p.wrote != ""
	}
	e.sig = e.word.String()
	e.adn = strings.HasPrefix(last, adnPrefix)
	e.det = len(t.pieces) == 0 && determiner(e.word.Parts[0].Names)
	e.conn = slices.Contains(connNames, last)
	if b > 0 && v > 0 && tailClass(t.pieces)&predTails != 0 && slices.Contains(sc.ix.auxStems(), stem) {
		e.back = sc.ix.structure(b, nil) - sc.ix.structure(0, nil)
	}
	return e
}
