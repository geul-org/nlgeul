//ff:func feature=token type=helper control=sequence
//ff:what adnNoun — is a candidate a noun of the word table, alone with no marker, of more than one syllable, whose last syllable is the copula 이 with a modifying closing fused into it (copulaAdnForms: 인, 일), so that the same letters read as the noun before it with the copula's modifying form (감정인 as 감정+이+ㄴ, 시행일 as 시행+이+ㄹ)
package token

import (
	"slices"
	"unicode/utf8"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func adnNoun(n rule.Near) bool {
	if n.Kind != tableNoun || n.Marks != "" || utf8.RuneCountInString(n.Stem) <= shortKey {
		return false
	}
	last, _ := utf8.DecodeLastRuneInString(n.Stem)
	return slices.Contains(copulaAdnForms, string(last))
}
