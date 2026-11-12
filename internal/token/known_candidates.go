//ff:func feature=token type=rule control=iteration dimension=1
//ff:what knownCandidates — the ways a stem is a known word with what follows it, one for each kind a table gives it whose graph lets it stand (and each with that kind's verdict); none when no table has it, or no kind of it takes that tail; the kind's graph is told whether the stem's word is the whole of its unit (sole)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Without a word table only the closed-set words are known, under
// their own names, when their set takes what follows them (an
// auxiliary verb's stem does not stand alone, a determiner takes no
// slot or closing); 이, 그 and 저 before particles or the copula are
// named as the pronouns they are there; a word in two sets is read once
// for each. In the first layer a stem of the table is known whatever
// follows. From the second layer on each kind's graph judges the stem. A
// closed-set word is named by its set only when it stands as that set's
// kind: 이 before a particle stands as the table's noun, not as the
// determiner, and is named so.
func (ix *index) knownCandidates(stem string, tail int, sole bool) []kindCandidate {
	c := ix.words[stem]
	if ix.layers == layerNone {
		return ix.setCandidates(c, stem, tail)
	}
	w := ix.lex[stem]
	passed, vs := 0, make([]float64, len(kindOrder))
	for i, k := range kindOrder {
		if (ix.closed[stem]|w.kinds)&k == 0 {
			continue
		}
		vs[i] = 1.0
		if ix.layers >= layerKind {
			vs[i] = ix.stemVerdict(k, rule.Piece{Form: stem, Tail: tail, Count: countOf(w, k, i), Sole: sole})
		}
		if vs[i] > 0 {
			passed |= k
		}
	}
	if passed == 0 {
		return nil
	}
	return passedCandidates(c, passed, vs)
}
