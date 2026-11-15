//ff:func feature=token type=rule control=sequence
//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables, and 보다 as an auxiliary after a question form (auxAskedRule); and 에서·서 as the adverbial particle before the nominative (adverbialFirstRule); and 이/가 right before 되다 or 아니다 as the complement particle (complementRule); and 라고·이라고 right after a closing quotation mark as the particle of direct quotation (quotedRule); and 한 standing alone as the determiner unless the word before it is a complement of 하다 (determinerOneRule); and 한 after the modifying 는 as the noun 한(限) (limitOneRule); and the 어/아 form right before an auxiliary verb (linkBeforeAuxRule), 말다 as the auxiliary verb after 지·고 (malLinkedRule) and 보니·보면 as the auxiliary after 다·다가 (auxWhileRule); and a noun of one syllable with its particle written onto a nominal (gluedShortRule); and a noun of the word table that ends in 인 or 일 where no word it could modify follows (unmodifiedRule); and 이란·란 right after a closing sign as the topic particle (shutTopicRule)

//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables, and 보다 as an auxiliary after a question form (auxAskedRule); and 에서·서 as the adverbial particle before the nominative (adverbialFirstRule); and 이/가 right before 되다 or 아니다 as the complement particle (complementRule); and 라고·이라고 right after a closing quotation mark as the particle of direct quotation (quotedRule); and 이란·란 right after a closing sign as the topic particle (shutTopicRule)

//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables, and 보다 as an auxiliary after a question form (auxAskedRule); and 에서·서 as the adverbial particle before the nominative (adverbialFirstRule); and 이/가 right before 되다 or 아니다 as the complement particle (complementRule); and 라고·이라고 right after a closing quotation mark as the particle of direct quotation (quotedRule); and 야 that ends a sentence as the ending of the copula (copulaEndRule)
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func (ix *index) addSentence() {
	aux := parse.ClosedWords()[auxSet]
	ix.sentence = []sentenceRule{auxLinkedRule(aux), mainUnlinkedRule(aux), auxAskedRule(), adverbialFirstRule(), complementRule(), quotedRule(), determinerOneRule(), limitOneRule()}
	// The 어/아 form of a main verb right before an auxiliary verb.
	ix.sentence = append(ix.sentence, linkBeforeAuxRule(aux))
	// 말다 after -지 or -고.
	ix.sentence = append(ix.sentence, malLinkedRule())
	// 보다 as an auxiliary in -다(가) 보니 and -다(가) 보면.
	ix.sentence = append(ix.sentence, auxWhileRule())
	// a noun of one syllable with its particle written onto a nominal
	ix.sentence = append(ix.sentence, gluedShortRule())
	// a noun ending in 인 or 일 where nothing an adnominal modifies follows
	ix.sentence = append(ix.sentence, unmodifiedRule())
	// 이란·란 right after a closing sign
	ix.sentence = append(ix.sentence, shutTopicRule())
	ix.sentence = append(ix.sentence, copulaEndRule())
}
