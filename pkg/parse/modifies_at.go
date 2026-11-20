//ff:func feature=parse type=reader control=sequence
//ff:what modifiesAt — do the letters from b to k read as a modifier and those from k to j as a known word of what it modifies: a negative adverb (negAdverbs: 안, 못, 아니) before a verb or an adjective, a numeral determiner (numDeterminers: 한, 두) before a nominal, an adverb of the word table before an adverb, or the modifying form of a predicate (adnFormAt) before the dependent noun 것 or 거 (adnOnto)
package parse

import "slices"

// Measured on the answer sheet, the modifying form before any nominal
// read apart many new words that hold a syllable with ㄴ or ㄹ under it
// (62 more wrong, 8 more right), and before any dependent noun of the
// closed set new words in an old spelling (4 more wrong, 3 more right):
// only 것 and 거 are asked, the dependent nouns whose one-syllable
// rebuttal is lifted there too (adnOnto). The lists are asked first, and
// the letters after k are read only when the letters before it may be
// such a modifier; for the dependent noun the syllable at k must begin as
// 것 and 거 (or 꺼 written for 거) do (onsetOfAny, boundForms) before the
// letters on either side are read.
func (sc *scan) modifiesAt(b, k, j int) bool {
	left := sc.letters(b, k)
	if slices.Contains(negAdverbs, left) && sc.readsAs(k, j, func(r stemCandidate) bool { return r.kind&(kindVerb|kindAdj) != 0 }) {
		return true
	}
	if slices.Contains(numDeterminers, left) && sc.readsAs(k, j, func(r stemCandidate) bool { return r.kind&kindNominal != 0 }) {
		return true
	}
	if sc.ix.lex[left].kinds&kindAdverb != 0 && sc.readsAs(k, j, func(r stemCandidate) bool { return r.kind&kindAdverb != 0 }) {
		return true
	}
	return onsetOfAny(sc.s[k], boundForms) && sc.adnFormAt(b, k) && sc.readsAs(k, j, func(r stemCandidate) bool { return slices.Contains(adnOnto, r.stem) })
}
