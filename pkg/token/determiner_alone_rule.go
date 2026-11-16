//ff:func feature=token type=rule control=sequence
//ff:what determinerAloneRule — the rule of a sentence "a demonstrative standing alone is the determiner": the candidate that reads 이, 그 or 저 (the set 지시어 of the closed set) as the determiner of the closed set (관형사:이) is preferred, but not when it is written onto the word before it (unless what stands right before it is an opening sign), nor when the word right after it is not one a modifier may stand before (not a nominal, an order, a determiner, nor a modifying or naming form: an adverb, a predicate, a sign, or the end of the sentence; unless it is an opening sign); it waits for the word after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The dictionary gives 이, 그 and 저 both as demonstrative determiners
// (이 사람, 그 책) and as demonstrative pronouns, and a pronoun is used
// with a particle (이는, 그가); the word table counts the pronoun more
// often. A determiner takes no particle and modifies the nominal after
// it (school grammar: 관형사는 체언 앞에서 그 체언을 꾸민다), so the
// one standing alone with a word it may modify after it is the
// determiner. A determiner is written apart from the word before it
// (clause 2 of the orthography); an opening quotation mark or bracket
// is written onto the words it encloses, not onto the word before it,
// so one right after such a sign stands alone too, and one right before
// such a sign modifies the word the sign opens.
func determinerAloneRule() sentenceRule {
	g, claim := sentenceClaim(ruleDeterminerAlone, rule.Listed, listed("지시어", rule.ListSelfStem, pronouns))
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	unmodifiable := except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not)
	unmodifiable.Attacks(claim)
	except(g, rule.Listed, whyOpenAfter, listed("여는 기호", rule.ListNextKind, []string{openName})).Attacks(unmodifiable)
	is := func(n rule.Near) bool {
		return n.Kind == determinerLabel+":"+n.Stem && slices.Contains(pronouns, n.Stem)
	}
	return sentenceRule{name: ruleDeterminerAlone, graph: g, bonus: bonusSameLetters, about: is, waits: is}
}
