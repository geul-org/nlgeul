//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentencePrior — the index of a reader holds the rule of a sentence rulePriorNoun after the rule of a word of two kinds standing alone as the determiner (tableDeterminerRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentencePrior(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	table, prior := at(ruleTableDeterminer), at(rulePriorNoun)
	if table < 0 || prior <= table {
		t.Fatal(table, prior, len(ix.sentence))
	}
}
