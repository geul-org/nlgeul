//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loneOf — a whole unit read as the markers of tail t alone: a word with no stem, marked as attaching to the unit before, its verdict no higher than a new word's (its stem is not a word this unit shows) nor than its pieces', and its preference their sum less the cost of a word's structure with one piece that leaves no trace (the stem, which is in another unit)
package token

import "math"

func (sc *scan) loneOf(t tail, end int) step {
	e := step{end: end, word: Eojeol{End: end, Attach: true}, verdict: sc.ix.unknown}
	e.pref = -sc.ix.structure(0, append([]piece{{unseen: true}}, t.pieces...))
	for _, p := range t.pieces {
		hidden, shown := hiddenOf(p.text)
		e.word.Parts = append(e.word.Parts, Part{Text: shown, Hidden: hidden, Names: namesOf(p.names), Unseen: p.unseen, Written: p.wrote})
		e.verdict, e.pref = math.Min(e.verdict, p.verdict), e.pref+p.pref
		e.typo = e.typo || p.wrote != ""
	}
	e.sig = e.word.String()
	return e
}
