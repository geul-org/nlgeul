//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceAlone — the index of a reader holds the rules of a sentence of a word of two kinds standing alone (adverbAloneRule, tableDeterminerRule, interjectionAloneRule) after the rule of 이·그·저 standing alone (determinerAloneRule), in this order (other rules may come after them)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceAlone(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	det, adv, table, inter := at(ruleDeterminerAlone), at(ruleAdverbAlone), at(ruleTableDeterminer), at(ruleInterjectionAlone)
	if det < 0 || adv <= det || table <= adv || inter <= table {
		t.Fatal(det, adv, table, inter, len(ix.sentence))
	}
}
