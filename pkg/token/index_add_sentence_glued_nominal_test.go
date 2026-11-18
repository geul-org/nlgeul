//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceGluedNominal — the index of a reader holds the rule of a sentence ruleGluedNominal after the rule of a word of a nominal and an adverb with the suffix 들 (pluralNounRule)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceGluedNominal(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	plural, glued := at(rulePluralNoun), at(ruleGluedNominal)
	if plural < 0 || glued <= plural {
		t.Fatal(plural, glued, len(ix.sentence))
	}
}
