//ff:func feature=token type=helper control=sequence
//ff:what shortNominal — is a candidate a noun of one syllable with a particle or the copula 이 right after it (자+의, 상+의, 일+이+ㄴ): its stem is one syllable named as a noun of the word table, a dependent noun or a pronoun (nounNames), its first marker is a particle (particleKinds) or the copula, and the word's letters are longer than its stem (the particle is written out, not put under the stem's last syllable)
package token

import (
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func shortNominal(n rule.Near) bool {
	marks := rule.MarksOf(n)
	kind, _, _ := strings.Cut(n.Kind, ":")
	if (n.Kind != tableNoun && !slices.Contains(nounNames, kind)) || utf8.RuneCountInString(n.Stem) != shortKey || utf8.RuneCountInString(n.Text) <= shortKey || len(marks) == 0 {
		return false
	}
	first, _, _ := strings.Cut(marks[0], ":")
	return slices.Contains(particleKinds, first) || marks[0] == copulaName
}
