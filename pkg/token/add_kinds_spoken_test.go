//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddKindsSpoken — on a kind's graph a spoken form of the closed table (spoken) seen once stands as it does seen enough times, as the rebuttal that it was rarely seen does not hold against it; a one-syllable form not in the table stands lower seen once; and the rebuttal that it is one syllable long still holds against 거
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddKindsSpoken(t *testing.T) {
	ix := &index{}
	ix.addKinds()
	g := ix.kinds[kindNominal]
	for _, form := range spoken {
		once, _ := verdictOf(g, rule.Piece{Form: form, Tail: rule.TailCase, Count: 1}, "", false)
		enough, _ := verdictOf(g, rule.Piece{Form: form, Tail: rule.TailCase, Count: seenEnough}, "", false)
		if once != enough {
			t.Fatal(form, once, enough)
		}
	}
	once, _ := verdictOf(g, rule.Piece{Form: "솥", Tail: rule.TailCase, Count: 1}, "", false)
	enough, _ := verdictOf(g, rule.Piece{Form: "솥", Tail: rule.TailCase, Count: seenEnough}, "", false)
	if once >= enough {
		t.Fatal("a one-syllable form not in the table", once, enough)
	}
	if geo, _ := verdictOf(g, rule.Piece{Form: "거", Tail: rule.TailCase, Count: 1}, "", false); geo >= 1 {
		t.Fatal("the rebuttal that it is one syllable long", geo)
	}
}
