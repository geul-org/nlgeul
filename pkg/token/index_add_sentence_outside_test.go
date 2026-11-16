//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceOutside — the index of a reader holds the rule of a sentence ruleOutsideNoun after the rule of a demonstrative standing alone (determinerAloneRule) (other rules may come after it)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceOutside(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, out := at(ruleDeterminerAlone), at(ruleOutsideNoun)
	if det < 0 || out <= det {
		t.Fatal(det, out, len(ix.sentence))
	}
}
