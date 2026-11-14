//ff:func feature=token type=reader control=sequence
//ff:what TestGrownDet — a verb of the table right after a determiner grows only the kept candidate whose last word is not one (마시+다), and both when it is itself a modifying form (마시+는); a noun grows both (사과+를)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGrownDet(t *testing.T) {
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{det: true}}}
	da := []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}
	verb := wordsScan("어떤마시다", layerCost)
	verb.detEnds[2] = true
	if got := verb.grown(nil, kept, tail{start: 4, pieces: da}, 2, 5); len(got) != 1 || got[0].prev != kept[0] {
		t.Fatal("마시+다 grows only the candidate whose last word is no determiner", got)
	}
	turned := wordsScan("어떤마시는", layerCost)
	turned.detEnds[2] = true
	if got := turned.grown(nil, kept, tail{start: 4, pieces: []piece{{text: "는", names: "꾸밈:는", verdict: 1, tail: rule.TailPred}}}, 2, 5); len(got) != 2 || !turned.adnEnds[5] {
		t.Fatal("마시+는 grows both candidates", got)
	}
	noun := wordsScan("어떤사과를", layerCost)
	noun.detEnds[2] = true
	if got := noun.grown(nil, kept, tail{start: 4, pieces: []piece{{text: "를", names: "목적격:를", verdict: 1, tail: rule.TailCase}}}, 2, 5); len(got) != 2 {
		t.Fatal("사과+를 grows both candidates", got)
	}
}
