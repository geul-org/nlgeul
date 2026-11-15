//ff:func feature=token type=rule control=sequence
//ff:what copulaEndRule — the rule of a sentence "야 that ends a sentence is the ending of the copula": a candidate one of whose markers is the ending 야 (copulaEndNames: it follows the copula 이, written or dropped) is preferred when nothing but a full stop or a question mark stands after it in its sentence (sentenceEnds), over 야 read as the vocative or the auxiliary particle; it waits for the words after it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The last word of a sentence is its predicate. A vocative (철수야) is an
// independent word and the auxiliary particle 야 (너야 알겠지) sets a word
// before what is said of it, so neither ends a sentence that a full stop
// or a question mark closes; the ending 야 of the copula does (그거야.).
// An exclamation mark and a comma also follow a vocative (철수야!,
// 철수야, 가자), so they are not read as the end.
func copulaEndRule() sentenceRule {
	g, claim := sentenceClaim(ruleCopulaEnd, rule.SelfHas, listed("서술격 뒤 맺음 야", rule.ListSelfMarks, copulaEndNames))
	except(g, rule.Listed, whyNotEnded, listed("문장을 닫는 부호", rule.ListNextLast, sentenceEnds), not).Attacks(claim)
	except(g, rule.Listed, whyWordAfter, listed("줄기 없음", rule.ListNextStem, []string{""}), not).Attacks(claim)
	is := func(n rule.Near) bool { return hasMark(n, copulaEndNames) }
	return sentenceRule{name: ruleCopulaEnd, graph: g, bonus: bonusCopula, about: is, waits: is}
}
