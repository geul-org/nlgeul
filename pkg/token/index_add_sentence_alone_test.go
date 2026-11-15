//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceAlone — the rules of a sentence an index holds include ruleAlone, once, after ruleQuoted, with a graph and functions about and waits
package token

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceAlone(t *testing.T) {
	var ix index
	ix.addSentence()
	var names []string
	for _, s := range ix.sentence {
		names = append(names, s.name)
	}
	at := slices.Index(names, ruleAlone)
	if at < 0 || slices.Index(names[at+1:], ruleAlone) >= 0 || at <= slices.Index(names, ruleQuoted) {
		t.Fatal(names)
	}
	if s := ix.sentence[at]; s.graph == nil || s.about == nil || s.waits == nil {
		t.Fatalf("%+v", s)
	}
}
