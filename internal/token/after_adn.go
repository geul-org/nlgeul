//ff:func feature=token type=rule control=iteration dimension=1
//ff:what afterAdn — does the word from b with tail t not stand right after a word of its unit that modifies a nominal, one that ends in a modifying closing or, when det is set, a determiner: its stem is a word of the word table there, and the graph of the kind it stands as (every kind of it when no kind is given) turns it down when told that such a word is before it and whether the word is itself a modifying or a naming form (a verb's or an adjective's stem, or a word that only stands alone: what modifies a nominal is followed by one; the graph lets through the auxiliaries made of a dependent noun and 하다 or 싶다 after a modifying closing, and a word that is itself a modifying or a naming form); no word is asked where no candidate has such a word end at b, or below the layer in which the kinds' graphs judge a stem
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) afterAdn(b int, t tail, det bool, kind int) bool {
	ends := sc.adnEnds
	if det {
		ends = sc.detEnds
	}
	if !ends[b] || t.as != "" || sc.ix.layers < layerKind {
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
		if kind != 0 && k != kind {
			continue
		}
		facts := rule.Piece{Form: stem, Tail: after, Count: countOf(w, k, i)}
		if (sc.ix.closed[stem]|w.kinds)&k == 0 || sc.ix.stemVerdict(k, facts) <= 0 {
			continue
		}
		facts.Adn, facts.Det, facts.Turned = !det, det, namesForm(t.pieces)
		v := sc.ix.stemVerdict(k, facts)
		stands, barred = stands || v > 0, barred || v <= 0
	}
	return barred && !stands
}
