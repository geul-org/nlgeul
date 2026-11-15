//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddPicked — each form of pickedClosings has a graph of its own, named with the form, and the count of graphs and edges grows by one each; the graph of 를 takes the closing 지 before it and turns down 고, 다 and 는지 for the reason that the particle is not written onto them
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddPicked(t *testing.T) {
	ix := &index{}
	ix.addPicked()
	if len(ix.pickedAux) != len(pickedClosings) || ix.graphs != len(pickedClosings) || ix.edges != len(pickedClosings) {
		t.Fatal(len(ix.pickedAux), ix.graphs, ix.edges)
	}
	e := ix.pickedAux["를"]
	if e == nil || e.name != pickedAuxName+"를" {
		t.Fatal("the graph of 를", e)
	}
	if v, _ := verdictOf(e.g, rule.Piece{Form: "를", Before: jiName, HasStem: true, Next: -1}, "", false); v <= 0 {
		t.Fatal("지 takes 를", v)
	}
	for _, before := range []string{"맺음:고", "맺음:다", "맺음:는지"} {
		if v, why := verdictOf(e.g, rule.Piece{Form: "를", Before: before, HasStem: true, Next: -1}, "", true); v > 0 || !slices.Contains(why, whyNotAfter) {
			t.Errorf("%s takes 를: %v %v", before, v, why)
		}
	}
}
