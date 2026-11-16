//ff:func feature=token type=rule control=sequence
//ff:what TestIndexAddSentenceHaOne — the index of a reader holds the rule of a sentence ruleHaOne after the rule of 한 after the modifying 는 (limitOneRule) and after the rule of 이·그·저 standing alone (determinerAloneRule) (other rules may come after it)
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceHaOne(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	limit, alone, ha := at(ruleLimitOne), at(ruleDeterminerAlone), at(ruleHaOne)
	if limit < 0 || alone < 0 || ha <= limit || ha <= alone {
		t.Fatal(limit, alone, ha, len(ix.sentence))
	}
}
