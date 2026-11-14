//ff:func feature=token type=reader control=iteration dimension=1
//ff:what fixed — the fused forms with a fixed first word that end at j (이게 → 이것+(이)): each is a word (written wrong, as 껄 for 걸, when the form is a wrong way to write it: the letters and the kinds of the closed-set word come with the read) and the particle under it, judged as that particle is before the one that follows; a form read only right after a modifying closing is judged here by whether some candidate has such a word end before it, and keeps its graph and facts to be asked again on each candidate (again)
package token

import (
	"unicode/utf8"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// next is the kind of the particle to the right (-1 for none).
func (sc *scan) fixed(j, next int) []tail {
	var out []tail
	for i, c := range sc.ix.contracts {
		e := sc.ix.fixed[i]
		b := j - utf8.RuneCountInString(c.form)
		if b < 0 || matchAt(sc.s[:j], b, c.form) != j {
			continue
		}
		facts := rule.Piece{Form: c.form, HasStem: true, Slot: linkOf(labelOf(c.name), c.text, e.slot), Next: next, Adn: sc.adnEnds[b] || (sc.attach && b == 0)}
		v := sc.judge(e, b, j, facts)
		if v <= 0 {
			continue
		}
		pc := piece{c.text, c.name, v, sc.ix.pref(e.count, c.text), tailOf(e), false, ""}
		read := tail{start: j, pieces: []piece{pc}, only: c.form, shown: c.stem, label: e.name}
		if c.wrote != "" {
			read.wrote, read.kinds = c.wrote, sc.ix.closed[c.stem]
		}
		if c.adn {
			read.again, read.facts = e, &facts
		}
		out = append(out, read)
	}
	return out
}
