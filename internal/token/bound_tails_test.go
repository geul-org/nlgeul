//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestBoundTails — the readings kept at b are grown once more by the word of a bound stem with a word written wrong, read as the right form; not for a stem written right, a tail fused into the stem, a stem a table names, or a changed stem
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundTails(t *testing.T) {
	sc := typoScan("입력/역활을", layerCost)
	kept := []*path{{verdict: 1}}
	eul := []piece{{"을", "목적격:을", 1, -1, rule.TailCase, false, ""}}
	got := sc.boundTails(nil, kept, tail{start: 5, pieces: eul}, 0, 6)
	if len(got) != 1 || got[0].sig != "입력/역할[대비:/,오타정정:입력/역활]+을[목적격:을]" || !got[0].typo {
		t.Fatal(got)
	}
	for n, tl := range []tail{{start: 5, pieces: eul, base: '활'}, {start: 5, pieces: eul, shown: "x"}, {start: 5, pieces: eul, as: numName}, {start: 5, pieces: eul, irr: irrS}} {
		if got := sc.boundTails(nil, kept, tl, 0, 6); len(got) != 0 {
			t.Errorf("%d: %v", n, got)
		}
	}
	if got := typoScan("입력/출력을", layerCost).boundTails(nil, kept, tail{start: 5, pieces: eul}, 0, 6); len(got) != 0 {
		t.Fatal("a stem written right", got)
	}
}
