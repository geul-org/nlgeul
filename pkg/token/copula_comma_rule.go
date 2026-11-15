//ff:func feature=token type=rule control=sequence
//ff:what copulaCommaRule — the rule of a sentence "이나 right before a comma is the copula with the linking ending -나": a candidate one of whose markers is the copula 이 written out (copulaNames) and whose last piece is the ending 나 (contrastNames) is preferred when a comma stands right after it, over 이나 read as the auxiliary particle; it waits for the words after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The auxiliary particle 이나 (밥이나 먹자, 너나 나나) sets a word before
// what is said of it; a comma right after a word ends a clause, and of
// the same letters the clause's end is the copula 이 with the linking
// ending -나, which sets what comes after against it (볼 것이나, 그렇지
// 않다). After a vowel the particle 나 and the dropped copula are not
// told apart so: 누구나, 어디나 stand before a comma as particles.
func copulaCommaRule() sentenceRule {
	g, claim := sentenceClaim(ruleCopulaComma, rule.SelfHas, listed("서술격 이", rule.ListSelfMarks, copulaNames))
	except(g, rule.Listed, whyNotContrast, listed("맺음 나", rule.ListSelfLast, contrastNames), not).Attacks(claim)
	except(g, rule.Listed, whyNoComma, listed("쉼표", rule.ListNextLast, commaNames), not).Attacks(claim)
	is := func(n rule.Near) bool { return hasMark(n, copulaNames) && slices.Contains(contrastNames, n.Last) }
	return sentenceRule{name: ruleCopulaComma, graph: g, bonus: bonusCopula, about: is, waits: is}
}
