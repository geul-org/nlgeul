//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceArgument — the index of a reader holds the rule of a sentence ruleArgumentAdn after the rule of a demonstrative standing alone (determinerAloneRule) (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceArgument(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, arg := at(ruleDeterminerAlone), at(ruleArgumentAdn)
	if det < 0 || arg <= det {
		t.Fatal(det, arg, len(ix.sentence))
	}
}
