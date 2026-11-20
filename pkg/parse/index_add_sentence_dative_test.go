//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceDative — the index of a reader holds the rule of a sentence ruleDativeContract after the rule of a demonstrative standing alone (determinerAloneRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceDative(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, dative := at(ruleDeterminerAlone), at(ruleDativeContract)
	if det < 0 || dative <= det {
		t.Fatal(det, dative, len(ix.sentence))
	}
}
