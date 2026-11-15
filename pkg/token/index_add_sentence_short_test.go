//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceShort — the index of a reader holds the rules of a sentence ruleGluedShort and ruleUnmodified after ruleQuoted, in this order
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceShort(t *testing.T) {
	var ix index
	ix.addSentence()
	var names []string
	for _, s := range ix.sentence {
		names = append(names, s.name)
	}
	at := func(name string) int { return slices.Index(names, name) }
	if at(ruleQuoted) < 0 || at(ruleGluedShort) <= at(ruleQuoted) || at(ruleUnmodified) <= at(ruleGluedShort) {
		t.Fatal(names)
	}
}
