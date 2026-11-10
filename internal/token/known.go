//ff:func feature=token type=rule control=iteration dimension=1
//ff:what known — is a stem a known word with what follows it: its name and its verdict, the best among the kinds a table gives it (0 when no table has it, or no kind of it takes that tail)
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

// Without a word table only the closed-set words are known, under
// their own names, when their set takes what follows them (an
// auxiliary verb's stem does not stand alone, a determiner takes no
// slot or closing). In the first
// layer a stem of the table is known whatever follows. From the second
// layer on each kind's graph judges the stem.
func (ix *index) known(stem string, tail int) (string, float64) {
	c, closed := ix.words[stem]
	if ix.layers == layerNone && closed {
		v, _ := verdictOf(c.g, rule.Piece{Form: stem, Tail: tail}, "")
		return c.name, v
	}
	if ix.layers == layerNone {
		return "", 0
	}
	w := ix.lex[stem]
	var names []string
	best := 0.0
	for i, k := range kindOrder {
		v := 1.0
		if ix.layers >= layerKind {
			v = ix.stemVerdict(k, tail, countOf(w, k, i), stem)
		}
		if (ix.closed[stem]|w.kinds)&k != 0 && v > 0 {
			names, best = append(names, kindLabel[k]), max(best, v)
		}
	}
	if closed && names != nil {
		return c.name, best
	}
	return "열쇠:" + strings.Join(names, "|"), best
}
