//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceHa — the index of a reader holds the rule of a sentence ruleHaLinked after the rule of a demonstrative standing alone (determinerAloneRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceHa(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, ha := at(ruleDeterminerAlone), at(ruleHaLinked)
	if det < 0 || ha <= det {
		t.Fatal(det, ha, len(ix.sentence))
	}
}
