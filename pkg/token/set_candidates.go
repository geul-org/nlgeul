//ff:func feature=token type=rule control=iteration dimension=1
//ff:what setCandidates — a closed-set word with no word table: one candidate for each set it is in (a pronoun named so before particles or the copula), with the verdict of its graph, which is one for the word; none when the word is in no set or its set does not take what follows
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (ix *index) setCandidates(c *entry, stem string, tail int) []kindCandidate {
	if c == nil {
		return nil
	}
	v := ix.judged(c.g, rule.Piece{Form: stem, Tail: tail}, "")
	if v <= 0 {
		return nil
	}
	var out []kindCandidate
	for _, name := range strings.Split(pronounName(c.name, stem, tail), "|") {
		out = append(out, kindCandidate{0, name, v})
	}
	return out
}
