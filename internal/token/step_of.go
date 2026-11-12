//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stepOf — the word from b to end that tail t makes when its stem stands as the candidate given: the stem and the markers after it as the parts of a word, with its verdict (the least of its pieces' and the stem's; the new word's graph for a stem that does not stand), its preference, the kind of the stem, and whether it ends in a modifying closing, is a determiner, or ends in the linking closing 어/아
package token

import (
	"math"
	"slices"
	"strings"
)

func (sc *scan) stepOf(b, end int, t tail, r stemRead) step {
	e := step{end: end, word: Eojeol{Parts: []Part{{Text: r.stem, Stem: true}}, End: end}, verdict: r.v, kind: r.kind}
	if r.v > 0 {
		e.word.Parts[0].Name, e.word.Parts[0].Written = nameOf(r.name), r.wrote
	} else {
		e.verdict = sc.unlisted(b, end, r.stem, t)
	}
	e.pref = sc.ix.stemPref(r.stem, tailClass(t.pieces), r.kind, r.v > 0, sc.tokens(b, t.start)) - sc.ix.structure(b, t.pieces)
	if t.as != "" {
		e.pref = t.weight - sc.ix.structure(b, t.pieces)
	}
	last := ""
	for _, p := range t.pieces {
		hidden, shown := hiddenOf(p.text)
		e.word.Parts = append(e.word.Parts, Part{Text: shown, Hidden: hidden, Name: nameOf(p.names), Unseen: p.unseen, Written: p.wrote})
		e.verdict, e.pref, last = math.Min(e.verdict, p.verdict), e.pref+p.pref, p.names
	}
	e.sig = e.word.String()
	e.adn = strings.HasPrefix(last, adnPrefix)
	e.det = len(t.pieces) == 0 && determiner(e.word.Parts[0].Name)
	e.conn = slices.Contains(connNames, last)
	if b > 0 && r.v > 0 && tailClass(t.pieces)&predTails != 0 && slices.Contains(sc.ix.auxStems(), r.stem) {
		e.back = sc.ix.structure(b, nil) - sc.ix.structure(0, nil)
	}
	return e
}
