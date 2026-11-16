//ff:func feature=token type=rule control=sequence
//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables, and 보다 as an auxiliary after a question form (auxAskedRule); and 에서·서 as the adverbial particle before the nominative (adverbialFirstRule); and 이/가 right before 되다 or 아니다 as the complement particle (complementRule); and 라고·이라고 right after a closing quotation mark as the particle of direct quotation (quotedRule); and 한 after the modifying 는 as the noun 한(限) (limitOneRule); and the 어/아 form right before an auxiliary verb (linkBeforeAuxRule), 말다 as the auxiliary verb after 지·고 (malLinkedRule) and 보니·보면 as the auxiliary after 다·다가 (auxWhileRule); and a noun of one syllable with its particle written onto a nominal (gluedShortRule); and a noun of the word table that ends in 인 or 일 where no word it could modify follows (unmodifiedRule); and 이란·란 right after a closing sign as the topic particle (shutTopicRule); and 야 and 라고 at a sentence's end as the copula's ending (copulaEndRule), 라고 after a number as the particle of quotation (numberQuoteRule) and 이나 before a comma as the copula with 나 (copulaCommaRule); and the rules of dependent nouns: a modifying form right before one (boundBeforeRule), 수 or 줄 as the noun where no modifying form stands before it (nounUnmodifiedRule), a dependent noun right after a modifying form (boundAfterRule), the noun of a dependent noun's letters before a middle dot (nounListedRule) and at the head of a sentence (nounHeadedRule), and a noun right after a modifier (nounModifiedRule); and after a number: a year, month and day written with dots as a date (dateFirstRule), numbers joined with a slash as a division (fractionFirstRule) and 도 at a word's end as the unit (degreeUnitRule); and 안·못 right before a verb or an adjective as the negative adverb (negAdverbRule)
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func (ix *index) addSentence() {
	aux := parse.ClosedWords()[auxSet]
	ix.sentence = []sentenceRule{auxLinkedRule(aux), mainUnlinkedRule(aux), auxAskedRule(), adverbialFirstRule(), complementRule(), quotedRule(), limitOneRule()}
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
	ix.sentence = append(ix.sentence, numberQuoteRule())
	ix.sentence = append(ix.sentence, copulaCommaRule())
	// A modifying form right before a dependent noun (boundBeforeRule).
	ix.sentence = append(ix.sentence, boundBeforeRule(boundNext, boundNominal))
	// The noun 수 or 줄 where no modifying form stands before it (nounUnmodifiedRule).
	ix.sentence = append(ix.sentence, nounUnmodifiedRule(adnBound))
	// A dependent noun right after a modifying form (boundAfterRule).
	ix.sentence = append(ix.sentence, boundAfterRule(boundNominal))
	// The noun of a dependent noun's letters before a middle dot (nounListedRule).
	ix.sentence = append(ix.sentence, nounListedRule(parse.ClosedWords()[boundSet]))
	// The noun of a dependent noun's letters at the head of a sentence (nounHeadedRule).
	ix.sentence = append(ix.sentence, nounHeadedRule(parse.ClosedWords()[boundSet]))
	// A noun right after a modifier (nounModifiedRule).
	ix.sentence = append(ix.sentence, nounModifiedRule(modifiedSkip))
	// A year, month and day written with dots as a date, and numbers joined with a slash as a division.
	ix.sentence = append(ix.sentence, dateFirstRule(), fractionFirstRule())
	// 도 at the end of a word right after a number as the unit.
	ix.sentence = append(ix.sentence, degreeUnitRule())
	// 안 or 못 right before a verb or an adjective as the negative adverb (negAdverbRule).
	ix.sentence = append(ix.sentence, negAdverbRule())
}
