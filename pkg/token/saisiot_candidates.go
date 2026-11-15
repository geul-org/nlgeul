//ff:func feature=token type=rule control=sequence
//ff:what saisiotCandidates — the candidate of a stem two nominals of the word table make, bound by 사이시옷 (saisiotParts: 전셋값 = 전세 + ㅅ + 값): a nominal named saisiotName, its verdict that of the weaker of the two, the front judged as standing alone and the back with what follows the stem, each with no rebuttal "one syllable long"; none when the stem is not so made, or either does not stand as a nominal
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The rebuttal "one syllable long" is lifted for both words, as it is
// for a word that is the whole of its unit: the ㅅ that 제30항 puts there
// sets where the front ends and the back begins, and the tail where the
// back ends, so neither is a run of letters met by chance inside a
// longer one. With the rebuttal a back word of one syllable stands lower
// than a longer back word that takes the particle into it, which then
// comes first (전깃불이 as 전기 + 불이, not 전깃불 + 이).
func (ix *index) saisiotCandidates(stem string, tail int) []kindCandidate {
	front, back, ok := ix.saisiotParts(stem)
	if !ok {
		return nil
	}
	v := min(ix.nominalVerdict(front, rule.TailBare, true), ix.nominalVerdict(back, tail, true))
	if v <= 0 {
		return nil
	}
	return []kindCandidate{{kindNominal, saisiotName, v}}
}
