//ff:func feature=parse type=rule control=sequence
//ff:what copulaEndRule — the rule of a sentence "야 and 라고 that end a sentence are endings": a candidate whose last piece is the ending 야 (it follows the copula 이, written or dropped) or the ending 라고 (copulaEndNames) is preferred when nothing but a full stop or a question mark stands after it in its sentence (sentenceEnds), over 야 read as the vocative or the auxiliary particle and 이라고 read as the particle of quotation; it waits for the words after it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The last word of a sentence is its predicate. A vocative (철수야) is an
// independent word and the auxiliary particle 야 (너야 알겠지) sets a word
// before what is said of it, so neither ends a sentence that a full stop
// or a question mark closes; the ending 야 of the copula does (그거야.).
// An exclamation mark and a comma also follow a vocative (철수야!,
// 철수야, 가자), so they are not read as the end. Nor does the particle
// 이라고, which stands before the verb that quotes (일품이라고 한다); the
// ending -라고 does (일품이라고.).
func copulaEndRule() sentenceRule {
	g, claim := sentenceClaim(ruleCopulaEnd, rule.Listed, listed("맺음 야·라고", rule.ListSelfLast, copulaEndNames))
	except(g, rule.Listed, whyNotEnded, listed("문장을 닫는 부호", rule.ListNextLast, sentenceEnds), not).Attacks(claim)
	except(g, rule.Listed, whyWordAfter, listed("줄기 없음", rule.ListNextStem, []string{""}), not).Attacks(claim)
	is := func(n rule.Near) bool { return slices.Contains(copulaEndNames, n.Last) }
	return sentenceRule{name: ruleCopulaEnd, graph: g, bonus: bonusShortTaken, about: is, waits: is}
}
