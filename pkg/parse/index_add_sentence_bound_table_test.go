//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceBoundTable — the index of a reader holds the rule of a sentence ruleBoundTable after the rule of a demonstrative standing alone (determinerAloneRule), in this order (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceBoundTable(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, table := at(ruleDeterminerAlone), at(ruleBoundTable)
	if det < 0 || table <= det {
		t.Fatal(det, table, len(ix.sentence))
	}
}
