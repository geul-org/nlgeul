//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceCompound — the index of a reader holds the rule of a sentence ruleCompoundFront after the rule of a word of a nominal and an adverb standing alone (adverbAloneRule)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceCompound(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	adv, front := at(ruleAdverbAlone), at(ruleCompoundFront)
	if adv < 0 || front <= adv {
		t.Fatal(adv, front, len(ix.sentence))
	}
}
