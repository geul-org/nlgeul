//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceGluedBound — the index of a reader holds the rule of a sentence ruleGluedBound after the rule of a noun of one syllable written onto a nominal (gluedShortRule) and the rule of a dependent noun of the word table's name after a modifying form (boundTableRule), in this order (other rules may come after it)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceGluedBound(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	short, table, glued := at(ruleGluedShort), at(ruleBoundTable), at(ruleGluedBound)
	if short < 0 || table <= short || glued <= table {
		t.Fatal(short, table, glued, len(ix.sentence))
	}
}
