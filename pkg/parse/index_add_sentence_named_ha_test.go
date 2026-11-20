//ff:func feature=parse type=rule control=sequence
//ff:what TestIndexAddSentenceNamedHa — the index of a reader holds the rule of a sentence ruleNamedBeforeHa after the rule of the -어야 form before 하다 or 되다 (linkBeforeHaRule)
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceNamedHa(t *testing.T) {
	var ix index
	ix.addSentence()
	at := func(name string) int {
		return slices.IndexFunc(ix.sentence, func(s sentenceRule) bool { return s.name == name })
	}
	link, named := at(ruleLinkBeforeHa), at(ruleNamedBeforeHa)
	if link < 0 || named <= link {
		t.Fatal(link, named, len(ix.sentence))
	}
}
