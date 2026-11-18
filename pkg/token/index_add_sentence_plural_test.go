//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentencePlural — the index of a reader holds the rule of a sentence rulePluralNoun after the rule of a word of a nominal and an adverb with a nominal written onto it (compoundFrontRule)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentencePlural(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	front, plural := at(ruleCompoundFront), at(rulePluralNoun)
	if front < 0 || plural <= front {
		t.Fatal(front, plural, len(ix.sentence))
	}
}
