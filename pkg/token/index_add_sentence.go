//ff:func feature=token type=rule control=sequence
//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func (ix *index) addSentence() {
	aux := parse.ClosedWords()[auxSet]
	ix.sentence = []sentenceRule{auxLinkedRule(aux), mainUnlinkedRule(aux)}
}
