//ff:func feature=token type=reader control=iteration dimension=1
//ff:what nounCopula — one predicate tail read after a dependent noun the copula's stem is fused into (nounCopulas): the tail's closing stands right after the noun's letter (테니) or is fused under it (텐데), the noun's graph judges it by whether some reading has a word ending in a modifying closing before it and by the closing, and the reading is the noun (터), the stem (이) and the tail, asked again on each reading (again); a tail that holds the copula or asks for a changed stem has no such reading
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The tail's stem is the one letter the noun is written as (테, or 텐
// with the closing's consonant under it), printed as the noun.
func (sc *scan) nounCopula(t tail) []tail {
	if t.close == nil || t.close.holds || t.irr != 0 || t.start < 1 {
		return nil
	}
	at := t.start - 1
	letter := sc.s[at]
	if t.base != 0 {
		letter = t.base
	}
	var out []tail
	for i, c := range nounCopulas {
		e := sc.ix.nouns[i]
		facts := rule.Piece{Form: c.form, HasStem: true, After: firstName(t), Adn: sc.adnEnds[at] || (sc.attach && at == 0)}
		if string(letter) != c.form || sc.judge(e, at, t.start, facts) <= 0 {
			continue
		}
		pc := piece{c.text, c.name, 1, sc.ix.pref(e.count, "이"), rule.TailCopula, false}
		out = append(out, tail{start: t.start, pieces: append([]piece{pc}, t.pieces...), noYo: t.noYo, close: t.close, only: string(sc.s[at:t.start]), shown: c.stem, label: e.name, again: e, facts: &facts})
	}
	return out
}
