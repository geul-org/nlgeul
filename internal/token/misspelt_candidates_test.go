//ff:func feature=token type=rule control=sequence
//ff:what TestMisspeltCandidates — the candidates of the right form of a misspelt stem, one for each kind both have whose graph lets it stand with what follows; none when no kind takes that tail or the two share no kind; every kind stands in the first layer; a closed-set word is named by its set
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMisspeltCandidates(t *testing.T) {
	ix := typoScan("x", layerCost).ix
	got := ix.misspeltCandidates("맡기", kindVerb, rule.TailPred, false)
	if len(got) != 1 || got[0].kind != kindVerb || got[0].name != "열쇠:동사" || got[0].v <= 0 || got[0].v >= 1 {
		t.Fatal(got)
	}
	if got = ix.misspeltCandidates("맡기", kindVerb, rule.TailCase, false); got != nil {
		t.Fatal("a verb takes no particle", got)
	}
	if got = ix.misspeltCandidates("맡기", kindNominal, rule.TailPred, false); got != nil {
		t.Fatal("the two share no kind", got)
	}
	if got = ix.misspeltCandidates("물", kindNominal|kindVerb, rule.TailCase, false); len(got) != 1 || got[0].kind != kindNominal {
		t.Fatal("only the kinds that take the tail", got)
	}
	if got = ix.misspeltCandidates("물", kindNominal|kindVerb, rule.TailBare, true); len(got) == 0 {
		t.Fatal("a word of two kinds", got)
	}
	if got = ix.misspeltCandidates("거", kindNominal, rule.TailBare, true); len(got) != 1 || got[0].name != "의존명사:거" {
		t.Fatal("a closed-set word is named by its set", got)
	}
	first := typoScan("x", layerKnown).ix.misspeltCandidates("맡기", kindVerb, rule.TailCase, false)
	if len(first) != 1 || first[0].v != 1 {
		t.Fatal("in the first layer every kind stands", first)
	}
}
