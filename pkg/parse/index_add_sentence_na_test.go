//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceNa — the index of a reader holds the rule of a sentence ruleNaLinked after the rule of 하다 after -어야 or a -기 form (haLinkedRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceNa(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	ha, na := at(ruleHaLinked), at(ruleNaLinked)
	if ha < 0 || na <= ha {
		t.Fatal(ha, na, len(ix.sentence))
	}
}
