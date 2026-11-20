//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceGive — the index of a reader holds the rule of a sentence ruleGiveLinked after the rule of a demonstrative standing alone (determinerAloneRule) (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceGive(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, give := at(ruleDeterminerAlone), at(ruleGiveLinked)
	if det < 0 || give <= det {
		t.Fatal(det, give, len(ix.sentence))
	}
}
