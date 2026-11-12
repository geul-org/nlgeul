//ff:func feature=token type=rule control=iteration dimension=1
//ff:what afterAdn — does the word from b with tail t not stand right after a word of its unit that ends in a modifying closing: its stem is a word of the word table there, and the graph of every kind it stands as turns it down when told that such a word is before it (a verb's or an adjective's stem: a modifying form modifies a nominal; the auxiliaries made of a dependent noun and 하다 or 싶다 are let through: 듯하, 만하); a word that is itself a modifying or a naming form is not asked (it modifies the same nominal, or stands as one), nor is any word where no reading has such a word end at b, or below the layer in which the kinds' graphs judge a stem
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) afterAdn(b int, t tail) bool {
	if !sc.adnEnds[b] || t.as != "" || sc.ix.layers < layerKind || namesForm(t.pieces) {
		return false
	}
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.shown != "" {
		stem = t.shown
	}
	w, after := sc.ix.lex[stem], tailClass(t.pieces)
	stands, barred := false, false
	for i, k := range kindOrder {
		count := countOf(w, k, i)
		if (sc.ix.closed[stem]|w.kinds)&k == 0 || sc.ix.stemVerdict(k, after, count, stem) <= 0 {
			continue
		}
		v, _ := verdictOf(sc.ix.kinds[k], rule.Piece{Form: stem, Tail: after, Count: count, Adn: true}, "")
		stands, barred = stands || v > 0, barred || v <= 0
	}
	return barred && !stands
}
