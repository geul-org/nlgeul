//ff:func feature=token type=reader control=iteration dimension=1
//ff:what fixed — the fused forms with a fixed first word that end at j (이게 → 이것+(이)): each is a word and the particle under it, judged as that particle is before the one that follows
package token

import "github.com/geul-org/koparse/pkg/rule"

// next is the kind of the particle to the right (-1 for none).
func (sc *scan) fixed(j, next int) []tail {
	var out []tail
	for i, c := range sc.ix.contracts {
		e := sc.ix.fixed[i]
		b := j - len([]rune(c.form))
		if b < 0 || string(sc.s[b:j]) != c.form {
			continue
		}
		v := sc.judge(e, b, j, rule.Piece{Form: c.form, HasStem: true, Slot: linkOf(labelOf(c.name), c.text, e.slot), Next: next, Adn: sc.adnEnds[b] || (sc.attach && b == 0)})
		if v <= 0 {
			continue
		}
		pc := piece{c.text, c.name, v, sc.ix.pref(e.count, c.text), tailOf(e), false}
		out = append(out, tail{start: j, pieces: []piece{pc}, only: c.form, shown: c.stem, label: e.name, adn: c.adn})
	}
	return out
}
