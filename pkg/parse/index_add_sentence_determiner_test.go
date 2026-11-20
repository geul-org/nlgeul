//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceDeterminer — the index of a reader holds the rule of a sentence ruleDeterminerAlone after the rule of a noun after a modifier (nounModifiedRule) and the rule of 도 after a number (degreeUnitRule), in this order (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceDeterminer(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	modified, degree, det := at(ruleNounModified), at(ruleDegreeUnit), at(ruleDeterminerAlone)
	if modified < 0 || degree <= modified || det <= degree {
		t.Fatal(modified, degree, det, len(ix.sentence))
	}
}
