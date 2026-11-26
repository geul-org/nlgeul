//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentencePossessive — the index of a reader holds the rule of a sentence rulePossessive after the rule of a word of a nominal and an interjection standing alone (interjectionAloneRule), and the rule of 내 read as 나 with 의 fused (rulePossessiveContract) right after it
package parse

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
	inter, owner, fused := at(ruleInterjectionAlone), at(rulePossessive), at(rulePossessiveContract)
	if inter < 0 || owner <= inter || fused != owner+1 {
		t.Fatal(inter, owner, fused, len(ix.sentence))
	}
}
