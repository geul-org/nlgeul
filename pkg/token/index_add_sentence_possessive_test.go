//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentencePossessive — the index of a reader holds the rule of a sentence rulePossessive after the rule of a word of a nominal and an interjection standing alone (interjectionAloneRule)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentencePossessive(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	inter, owner := at(ruleInterjectionAlone), at(rulePossessive)
	if inter < 0 || owner <= inter {
		t.Fatal(inter, owner, len(ix.sentence))
	}
}
