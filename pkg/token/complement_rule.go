//ff:func feature=token type=rule control=sequence
//ff:what complementRule — the rule of a sentence "이/가 right before 되다 or 아니다 is the complement particle": a candidate one of whose markers is the complement 이 or 가 (complementNames) is preferred when the word right after it is 되다 or 아니다 (complementStems); when the word right after it is a question word with 이/가 and 되다 or 아니다 comes after that, the candidate before the question word is preferred and the question word itself is not (반장이 누가 됐어: 반장이 is the complement)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The candidate waits for the words after it up to the predicate.
func complementRule() sentenceRule {
	g, claim := sentenceClaim(ruleComplement, rule.SelfHas, listed("보격 이·가", rule.ListSelfMarks, complementNames))
	far := except(g, rule.Listed, whyNoBecome, listed("되다·아니다", rule.ListNextStem, complementStems), not)
	far.Attacks(claim)
	asks := except(g, rule.Listed, whyAsksNext, listed("의문 대명사", rule.ListNextStem, askWords))
	asks.Attacks(far)
	except(g, rule.Listed, whyNoBecomeAfter, listed("되다·아니다", rule.ListNext2Stem, complementStems), not).Attacks(asks)
	except(g, rule.Listed, whyAsksBare, listed("이·가", rule.ListNextLast, subjectNames), not).Attacks(asks)
	self := except(g, rule.Listed, whyAsksSelf, listed("의문 대명사", rule.ListSelfStem, askWords))
	self.Attacks(claim)
	except(g, rule.Listed, whyNoSubjectBefore, listed("이·가", rule.ListPrevLast, subjectNames), not).Attacks(self)
	is := func(n rule.Near) bool { return hasMark(n, complementNames) }
	return sentenceRule{name: ruleComplement, graph: g, bonus: bonusSameLetters, about: is, waits: is}
}
