//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceNominalBound — the index of a reader holds the rule of a sentence ruleNominalBound after the rule of a dependent noun of the word table's name after a modifying form (boundTableRule), in this order (other rules may come after it)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceNominalBound(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	table, nominal := at(ruleBoundTable), at(ruleNominalBound)
	if table < 0 || nominal <= table {
		t.Fatal(table, nominal, len(ix.sentence))
	}
}
