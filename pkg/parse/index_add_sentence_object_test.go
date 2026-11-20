//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceObject — the index of a reader holds the rule of a sentence rulePredAfterObject after the rule of a demonstrative standing alone (determinerAloneRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceObject(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, object := at(ruleDeterminerAlone), at(rulePredAfterObject)
	if det < 0 || object <= det {
		t.Fatal(det, object, len(ix.sentence))
	}
}
