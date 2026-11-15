//ff:func feature=token type=rule control=sequence
//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables, and 보다 as an auxiliary after a question form (auxAskedRule); and 에서·서 as the adverbial particle before the nominative (adverbialFirstRule); and 이/가 right before 되다 or 아니다 as the complement particle (complementRule); and 라고·이라고 right after a closing quotation mark as the particle of direct quotation (quotedRule); and 한 standing alone as the determiner unless the word before it is a complement of 하다 (determinerOneRule); and 한 after the modifying 는 as the noun 한(限) (limitOneRule)
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func (ix *index) addSentence() {
	aux := parse.ClosedWords()[auxSet]
	ix.sentence = []sentenceRule{auxLinkedRule(aux), mainUnlinkedRule(aux), auxAskedRule(), adverbialFirstRule(), complementRule(), quotedRule(), determinerOneRule(), limitOneRule()}
}
