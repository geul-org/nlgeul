//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceNeg — the index of a reader holds the rule of a sentence ruleLinkBeforeNeg after the rule of a demonstrative standing alone (determinerAloneRule)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceNeg(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, neg := at(ruleDeterminerAlone), at(ruleLinkBeforeNeg)
	if det < 0 || neg <= det {
		t.Fatal(det, neg, len(ix.sentence))
	}
}
