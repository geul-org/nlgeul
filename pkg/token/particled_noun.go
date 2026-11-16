//ff:func feature=token type=helper control=sequence
//ff:what particledNoun — is a candidate a noun of the word table (열쇠:체언) whose stem is one of words and whose first marker is a particle (particleKinds), not the copula (측+은, 자+가; not 내+(이)+지)
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func particledNoun(n rule.Near, words []string) bool {
	marks := rule.MarksOf(n)
	if n.Kind != tableNoun || len(marks) == 0 || !slices.Contains(words, n.Stem) {
		return false
	}
	first, _, _ := strings.Cut(marks[0], ":")
	return slices.Contains(particleKinds, first)
}
