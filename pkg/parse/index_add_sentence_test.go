//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestIndexAddSentence — the index of a reader holds the rules of a sentence with the names ruleAuxLinked and ruleMainUnlinked in this order (other rules may be among them), each with a graph and a function about; and the stream of a reader carries them
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentence(t *testing.T) {
	r := streamReader(t, "words-aux-linked.tsv")
	var names []string
	for _, s := range r.ix.sentence {
		names = append(names, s.name)
		if s.graph == nil || s.about == nil {
			t.Fatalf("%+v", s)
		}
	}
	at := func(name string) int { return slices.Index(names, name) }
	if at(ruleAuxLinked) < 0 || at(ruleAuxLinked) >= at(ruleMainUnlinked) {
		t.Fatal(names)
	}
	if got := len(r.Stream("있다").rules); got != len(names) {
		t.Fatal("the stream takes the index's rules", got, len(names))
	}
	var ix index
	ix.addSentence()
	if len(ix.sentence) != len(names) {
		t.Fatal(len(ix.sentence), len(names))
	}
}
