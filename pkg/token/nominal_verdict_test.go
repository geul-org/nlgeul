//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNominalVerdict — the verdict of a stem's nominal candidate: 값 before a case particle stands as the table's nominal, higher as the whole of its unit (no rebuttal "one syllable long") than inside one; 비, a noun and a verb, has its noun's verdict, not the verb's; 것 of the closed set and of the table has the verdict of its candidates; 0 for a verb (앉), for a nominal before a predicate tail and for a stem no table lists
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNominalVerdict(t *testing.T) {
	ix := streamReader(t, "words-saisiot.tsv").ix
	whole, inside := ix.nominalVerdict("값", rule.TailCase, true), ix.nominalVerdict("값", rule.TailCase, false)
	if whole <= inside || inside <= 0 {
		t.Fatal(whole, inside)
	}
	w := ix.lex["값"]
	if want := ix.stemVerdict(kindNominal, rule.Piece{Form: "값", Tail: rule.TailCase, Count: countOf(w, kindNominal, 0), Sole: true}); whole != want {
		t.Fatal("the nominal graph's verdict", whole, want)
	}
	noun := ix.nominalVerdict("비", rule.TailBare, false)
	if got := ix.knownCandidates("비", rule.TailBare, false, false, false, false); len(got) != 1 || got[0].kind != kindNominal || got[0].v != noun || noun <= 0 {
		t.Fatal("the noun's verdict", got, noun)
	}
	if v := ix.nominalVerdict("것", rule.TailCase, true); v <= 0 || v != ix.known("것", rule.TailCase, true) {
		t.Fatal("것", v)
	}
	for _, c := range []struct {
		stem string
		tail int
	}{{"앉", rule.TailBare}, {"값", rule.TailPred}, {"없는말", rule.TailCase}} {
		if v := ix.nominalVerdict(c.stem, c.tail, true); v != 0 {
			t.Error(c.stem, v)
		}
	}
}
