//ff:func feature=token type=rule control=iteration dimension=1
//ff:what misspeltCandidates — the ways the right form of a misspelt stem is a known word with what follows it, one for each kind both the right form (in the word table or as a closed-set word) and the misspelt form have, whose graph lets it stand; none when no kind takes that tail; each kind's graph is told that the stem is a misspelt form; a closed-set word is named by its set when it stands as that set's kind
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// It is knownCandidates's second layer for a misspelt stem: in the first
// layer every kind counts as standing. The right form's own count is the
// stem's (the count of a misspelt line is 0, nothing was counted).
// A right form that is a closed-set word has the kind of its set too,
// and is named by the set as knownCandidates names it.
func (ix *index) misspeltCandidates(right string, kinds, tail int, sole bool) []kindCandidate {
	w := ix.lex[right]
	have := ix.closed[right] | w.kinds
	passed, vs := 0, make([]float64, len(kindOrder))
	for i, k := range kindOrder {
		if have&kinds&k == 0 {
			continue
		}
		vs[i] = 1.0
		if ix.layers >= layerKind {
			vs[i] = ix.stemVerdict(k, rule.Piece{Form: right, Tail: tail, Count: countOf(w, k, i), Sole: sole, Misspelt: true})
		}
		if vs[i] > 0 {
			passed |= k
		}
	}
	if passed == 0 {
		return nil
	}
	return passedCandidates(ix.words[right], passed, vs)
}
