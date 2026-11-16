//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceApart — the index of a reader holds the rule of a sentence ruleApartParticle after the rule of a demonstrative standing alone (determinerAloneRule) (other rules may come after it)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceApart(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, apart := at(ruleDeterminerAlone), at(ruleApartParticle)
	if det < 0 || apart <= det {
		t.Fatal(det, apart, len(ix.sentence))
	}
}
