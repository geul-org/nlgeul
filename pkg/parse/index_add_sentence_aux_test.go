//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestIndexAddSentenceAux — the index of a reader holds the rules ruleLinkBeforeAux and ruleMalLinked once each, after ruleMainUnlinked; of them only ruleLinkBeforeAux waits for the words after a candidate
package parse

import (
	"slices"
	"testing"
)

func TestIndexAddSentenceAux(t *testing.T) {
	r := streamReader(t, "words-link-before-aux.tsv")
	var names []string
	waits := map[string]bool{}
	for _, s := range r.ix.sentence {
		names = append(names, s.name)
		waits[s.name] = s.waits != nil
	}
	main := slices.Index(names, ruleMainUnlinked)
	for _, name := range []string{ruleLinkBeforeAux, ruleMalLinked} {
		at := slices.Index(names, name)
		if at <= main || slices.Index(names[at+1:], name) >= 0 {
			t.Fatal(name, names)
		}
		if waits[name] != (name == ruleLinkBeforeAux) {
			t.Fatal(name, "waits", waits[name])
		}
	}
}
