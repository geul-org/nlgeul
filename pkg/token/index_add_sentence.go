//ff:func feature=token type=rule control=sequence
//ff:what addSentence — the rules that judge a candidate with the words of its sentence around it, in the order they are asked: an auxiliary verb after a linking form (auxLinkedRule) and the main predicate elsewhere (mainUnlinkedRule), both of the closed set of auxiliary verbs of the grammar tables; and 에서·서 as the adverbial particle before the nominative (adverbialFirstRule); and 이/가 right before 되다 or 아니다 as the complement particle (complementRule); and 라고·이라고 right after a closing quotation mark as the particle of direct quotation (quotedRule); and 한 after the modifying 는 as the noun 한(限) (limitOneRule); and the 어/아 form right before an auxiliary verb (linkBeforeAuxRule) and 말다 as the auxiliary verb after 지·고 (malLinkedRule); and a noun of one syllable with its particle written onto a nominal (gluedShortRule); and a noun of the word table that ends in 인 or 일 where no word it could modify follows (unmodifiedRule); and 이란·란 right after a closing sign as the topic particle (shutTopicRule); and 야 and 라고 at a sentence's end as the copula's ending (copulaEndRule), 라고 after a number as the particle of quotation (numberQuoteRule) and 이나 before a comma as the copula with 나 (copulaCommaRule); and the rules of dependent nouns: a modifying form right before one (boundBeforeRule), 수 or 줄 as the noun where no modifying form stands before it (nounUnmodifiedRule), a dependent noun right after a modifying form (boundAfterRule), the noun of a dependent noun's letters at the head of a sentence or before a middle dot (nounHeadedRule), and a noun right after a modifier (nounModifiedRule), and the noun of a dependent noun's letters with neither a modifier nor a nominal right before it (nounUnheldRule); and after a number: 도 at a word's end as the unit (degreeUnitRule); and 이·그·저 standing alone, before a word it may modify, as the determiner (determinerAloneRule); and 이런·그런·저런·어떤·다른 standing alone as the determiner, but 다른 after 와/과 as the adjective's modifying form (determinerFormRule); and 한 after a complement of 하다 (the object particle, 로/으로, -게, -도록, -다고 …), or written onto a word, as 하다 with the modifying ㄴ (haOneRule); and a word the word table gives both as a nominal and as another kind, standing alone with no particle: as the adverb (adverbAloneRule), before a word it may modify as the determiner (tableDeterminerRule), and as the interjection (interjectionAloneRule); and 내, 네 and 제 standing alone before a word a modifier may stand before as the pronoun with 의 fused (possessiveRule);; and a word the word table gives both as a nominal and as an adverb with a nominal written onto it as the nominal, the front of a compound noun (compoundFrontRule);; and such a word with the suffix 들 after it as the nominal (pluralNounRule);; and a word the word table gives both as a nominal and as an adverb or a determiner, written onto a nominal before it, as the nominal (gluedNominalRule); and a dependent noun the word table names a noun, with its marker, right after a modifying form (boundTableRule); and such a dependent noun written onto a nominal with its particle (gluedBoundRule); and a dependent noun the word table names a noun that stands after a nominal, with its particle, right after a nominal (nominalBoundRule); and a noun of one syllable with its particle right after a numeral determiner of the closed set (countedShortRule); and a verb's or an adjective's modifying form between a word ending in a case particle of an argument and a word it may modify (argumentAdnRule); and 밖에 with no word of negation after it as the noun 밖 with the adverbial 에 (outsideNounRule); and 달다 after the 어/아 form as the auxiliary verb of asking (giveLinkedRule); and a word of markers written apart as the particle of the word before (apartParticleRule); and 내게, 네게 and 제게 as the pronoun with the particle 게 but before 하다, 되다 or 만들다 (dativeContractRule); and the -지 form of a predicate right before 않다, 못하다 or 말다 (linkBeforeNegRule); and a verb or an adjective, not in its modifying form, right after an object (predAfterObjectRule); and 하다 as the verb right after -어야 or a -기 form (haLinkedRule); and 전 after a naming form or a quantity as the noun (priorNounRule); and 나다 as the auxiliary verb right after -고 (naLinkedRule)
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func (ix *index) addSentence() {
	aux := parse.ClosedWords()[auxSet]
	ix.sentence = []sentenceRule{auxLinkedRule(aux), mainUnlinkedRule(aux), adverbialFirstRule(), complementRule(), quotedRule(), limitOneRule()}
	// The 어/아 form of a main verb right before an auxiliary verb.
	ix.sentence = append(ix.sentence, linkBeforeAuxRule(aux))
	// 말다 after -지 or -고.
	ix.sentence = append(ix.sentence, malLinkedRule())
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
	// The noun of a dependent noun's letters at the head of a sentence or before a middle dot (nounHeadedRule).
	ix.sentence = append(ix.sentence, nounHeadedRule(parse.ClosedWords()[boundSet]))
	// A noun right after a modifier (nounModifiedRule).
	ix.sentence = append(ix.sentence, nounModifiedRule(modifiedSkip))
	// The noun of a dependent noun's letters with no modifier and no nominal before it (nounUnheldRule).
	ix.sentence = append(ix.sentence, nounUnheldRule(unheldBound))
	// 도 at the end of a word right after a number as the unit.
	ix.sentence = append(ix.sentence, degreeUnitRule())
	// 이·그·저 standing alone as the determiner (determinerAloneRule).
	ix.sentence = append(ix.sentence, determinerAloneRule())
	// 이런·그런·저런·어떤·다른 standing alone as the determiner (determinerFormRule).
	ix.sentence = append(ix.sentence, determinerFormRule())
	// 한 after a complement of 하다 as 하다 with the modifying ㄴ (haOneRule).
	ix.sentence = append(ix.sentence, haOneRule())
	// A word of two kinds (a nominal and an adverb) standing alone as the adverb (adverbAloneRule).
	ix.sentence = append(ix.sentence, adverbAloneRule(ix.lex))
	// A word of two kinds (a nominal and a determiner) standing alone before a word it may modify as the determiner (tableDeterminerRule).
	ix.sentence = append(ix.sentence, tableDeterminerRule(ix.lex, ix.closed))
	// A word of two kinds (a nominal and an interjection) standing alone as the interjection (interjectionAloneRule).
	ix.sentence = append(ix.sentence, interjectionAloneRule(ix.lex, ix.closed))
	// 내, 네 and 제 standing alone before a word a modifier may stand before as the pronoun with 의 fused (possessiveRule).
	ix.sentence = append(ix.sentence, possessiveRule())
	// A word of two kinds (a nominal and an adverb) with a nominal written onto it as the nominal (compoundFrontRule).
	ix.sentence = append(ix.sentence, compoundFrontRule(ix.lex))
	// A word of two kinds (a nominal and an adverb) with the suffix 들 as the nominal (pluralNounRule).
	ix.sentence = append(ix.sentence, pluralNounRule(ix.lex))
	// A word of two kinds (a nominal and an adverb or a determiner) written onto a nominal as the nominal (gluedNominalRule).
	ix.sentence = append(ix.sentence, gluedNominalRule(ix.lex, ix.closed))
	// A dependent noun of the word table's name right after a modifying form (boundTableRule).
	ix.sentence = append(ix.sentence, boundTableRule(tableBound))
	// A dependent noun of the word table's name written onto a nominal with its particle (gluedBoundRule).
	ix.sentence = append(ix.sentence, gluedBoundRule(tableBound))
	// A dependent noun of the word table's name after a nominal with its particle (nominalBoundRule).
	ix.sentence = append(ix.sentence, nominalBoundRule(nominalBound))
	// A noun of one syllable with its particle after a numeral determiner (countedShortRule).
	ix.sentence = append(ix.sentence, countedShortRule(numDeterminers))
	// A modifying form between an argument and a word it may modify (argumentAdnRule).
	ix.sentence = append(ix.sentence, argumentAdnRule())
	// 밖에 with no negative after it as the noun 밖 with 에 (outsideNounRule).
	ix.sentence = append(ix.sentence, outsideNounRule())
	// 달다 after the 어/아 form as the auxiliary verb of asking (giveLinkedRule).
	ix.sentence = append(ix.sentence, giveLinkedRule())
	// Markers written apart as the particle of the word before (apartParticleRule).
	ix.sentence = append(ix.sentence, apartParticleRule())
	// 내게, 네게 and 제게 as the pronoun with the particle 게 (dativeContractRule).
	ix.sentence = append(ix.sentence, dativeContractRule())
	// The -지 form right before 않다, 못하다 or 말다 (linkBeforeNegRule).
	ix.sentence = append(ix.sentence, linkBeforeNegRule())
	// A predicate right after an object (predAfterObjectRule).
	ix.sentence = append(ix.sentence, predAfterObjectRule())
	// 하다 right after -어야 or a -기 form as the verb (haLinkedRule).
	ix.sentence = append(ix.sentence, haLinkedRule())
	// 전 right after a naming form or a quantity as the noun (priorNounRule).
	ix.sentence = append(ix.sentence, priorNounRule(ix.lex))
	// 나다 right after -고 as the auxiliary verb (naLinkedRule).
	ix.sentence = append(ix.sentence, naLinkedRule())
}
