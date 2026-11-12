//ff:func feature=token type=rule control=iteration dimension=1
//ff:what setReads — a closed-set word with no word table: one reading for each set it is in (a pronoun named so before particles or the copula), with the verdict of its graph, which is one for the word; none when the word is in no set or its set does not take what follows
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (ix *index) setReads(c *entry, stem string, tail int) []kindRead {
	if c == nil {
		return nil
	}
	v := verdict(c.g, rule.Piece{Form: stem, Tail: tail})
	if v <= 0 {
		return nil
	}
	var out []kindRead
	for _, name := range strings.Split(pronounName(c.name, stem, tail), "|") {
		out = append(out, kindRead{0, name, v})
	}
	return out
}
