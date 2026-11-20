//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceLinkHa — the index of a reader holds the rule of a sentence ruleLinkBeforeHa after the rule of 하다 after -어야 or a -기 form (haLinkedRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceLinkHa(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	ha, link := at(ruleHaLinked), at(ruleLinkBeforeHa)
	if ha < 0 || link <= ha {
		t.Fatal(ha, link, len(ix.sentence))
	}
}
