//ff:func feature=parse type=rule control=sequence
//ff:what determinerFormRule — the rule of a sentence "a demonstrative determiner that is also an adjective's modifying form, standing alone, is the determiner": the candidate that reads 이런, 그런, 저런, 어떤 or 다른 (formDeterminers) as a determiner (열쇠:관형사, or 관형사:<word> of the closed set) is preferred, by bonusOverUncounted, but not when it is written onto the word before it (unless what stands right before it is an opening sign), nor, for 다른 alone, when the word before it ends in 와/과, the comparison 다르다 takes (comparedNames), or has it with a particle more (과는: rule.PrevHas), where it is the adjective's modifying form
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The dictionary gives 이런, 그런, 저런, 어떤 and 다른 as determiners
// (그런 「관형사」 상태, 모양, 성질 따위가 그러한; 다른 「관형사」 당장
// 문제 되거나 해당되는 것 이외의), and school grammar counts them among
// the demonstrative determiners; their letters are also the modifying
// form of the adjectives 이렇다, 그렇다, 저렇다, 어떻다 and 다르다. The
// word table counts the adjectives about two million times and these
// determiners about once, so a determiner standing alone loses to the
// adjective's form. A determiner takes no argument: 다르다 takes the
// thing it is compared with in …와/…과 (비교가 되는 두 대상이 서로 같지
// 아니하다), so 다른 right after a word ending in 와/과 is the adjective's
// form. A determiner is written apart from the word before it (clause 2
// of the orthography), and an opening sign is written onto the words it
// encloses. A particle after a case particle adds a sense and does not
// change what the word is to the predicate (…과는 다른).
func determinerFormRule() sentenceRule {
	g, claim := sentenceClaim(ruleDeterminerForm, rule.Listed, listed("지시 관형사", rule.ListSelfStem, formDeterminers))
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	compared := except(g, rule.PrevHas, whyComparedBefore, listed("와·과", rule.ListPrevLast, comparedNames))
	compared.Attacks(claim)
	except(g, rule.Listed, whyNotDiffer, listed("다른", rule.ListSelfStem, []string{differStem}), not).Attacks(compared)
	is := func(n rule.Near) bool {
		return slices.Contains(formDeterminers, n.Stem) && (n.Kind == keyKind+":"+determinerLabel || n.Kind == determinerLabel+":"+n.Stem)
	}
	return sentenceRule{name: ruleDeterminerForm, graph: g, bonus: bonusOverUncounted, about: is}
}
