//ff:func feature=token type=rule control=iteration dimension=1
//ff:what known — is a stem a known word with what follows it: its name and its verdict, the best among the kinds a table gives it (0 when no table has it, or no kind of it takes that tail)
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Without a word table only the closed-set words are known, under
// their own names, when their set takes what follows them (an
// auxiliary verb's stem does not stand alone, a determiner takes no
// slot or closing); 이, 그 and 저 before particles or the copula are
// named as the pronouns they are there. In the first
// layer a stem of the table is known whatever follows. From the second
// layer on each kind's graph judges the stem. A closed-set word is
// named by its set only when it stands as that set's kind: 이 before a
// particle stands as the table's noun, not as the determiner, and is
// named so.
func (ix *index) known(stem string, tail int) (string, float64) {
	c, closed := ix.words[stem]
	if ix.layers == layerNone && closed {
		return pronounName(c.name, stem, tail), verdict(c.g, rule.Piece{Form: stem, Tail: tail})
	}
	if ix.layers == layerNone {
		return "", 0
	}
	w := ix.lex[stem]
	var names []string
	best, passed := 0.0, 0
	for i, k := range kindOrder {
		v := 1.0
		if (ix.closed[stem]|w.kinds)&k != 0 && ix.layers >= layerKind {
			v = ix.stemVerdict(k, rule.Piece{Form: stem, Tail: tail, Count: countOf(w, k, i)})
		}
		if (ix.closed[stem]|w.kinds)&k != 0 && v > 0 {
			names, best, passed = append(names, kindLabel[k]), max(best, v), passed|k
		}
	}
	if own := ownName(c, passed); closed && own != "" {
		return own, best
	}
	return "열쇠:" + strings.Join(names, "|"), best
}
