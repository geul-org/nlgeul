//ff:func feature=token type=reader control=sequence
//ff:what TestAddTypo — the entry of a closing written wrong is put under its key with a graph of one edge more than the right closing's (the rebuttal of a misspelt form), which stands the closing below the right one when it is told the closing is misspelt; its count is not added to the total, its form is not put into the set of closings, and the entry already there is not changed
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddTypo(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}, closingSet: map[string]bool{}}
	right := &entry{name: "맺음:ㄹ걸", form: "걸", class: classClosing, lead: rule.JongL, after: rule.AfterAny, count: 40}
	ix.add(right)
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	wrong := typoClosingEntry(right, typoClosing{"껄", rule.JongL, "맺음:ㄹ걸"})
	ix.addTypo(wrong)
	if ix.graphs != graphs+1 || ix.edges != edges+edges+1 || ix.total != total || len(ix.closingSet) != 0 {
		t.Fatal(ix.graphs-graphs, ix.edges-edges, ix.total-total, ix.closingSet)
	}
	if len(ix.byKey['껄']) != 1 || ix.byKey['껄'][0] != wrong || wrong.g == nil || wrong.g == right.g {
		t.Fatal(ix.byKey)
	}
	if right.wrote != "" || len(ix.byKey['걸']) != 1 || ix.byKey['걸'][0] != right {
		t.Fatal("the right closing stays as it is", ix.byKey['걸'])
	}
	p := rule.Piece{HasStem: true, Next: -1, Base: "하", Form: "걸"}
	v, _ := verdictOf(right.g, p, "할걸", false)
	p.Misspelt = true
	typo, _ := verdictOf(wrong.g, p, "할껄", false)
	if v != 1 || math.Abs(typo-(2/(1+forceTypo)-1)) > 1e-9 {
		t.Fatal("a closing written wrong stands below the right one by the force of a misspelt form", v, typo)
	}
	p.Misspelt = false
	if untold, _ := verdictOf(wrong.g, p, "할껄", false); untold != 1 {
		t.Fatal("the rebuttal is for the facts that say the closing is misspelt", untold)
	}
}
