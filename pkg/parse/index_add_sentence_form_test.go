//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceForm — the index of a reader holds the rule of a sentence ruleDeterminerForm after the rule of 이·그·저 standing alone (determinerAloneRule), in this order (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceForm(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	alone, form := at(ruleDeterminerAlone), at(ruleDeterminerForm)
	if alone < 0 || form <= alone {
		t.Fatal(alone, form, len(ix.sentence))
	}
}
