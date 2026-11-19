//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddChains — each ending of closingChains has a graph of its own, named with the ending's name, and the count of graphs and edges grows by one each; the graph of 거나 takes the closings 다 and ㄴ다 before it and turns down 고, 지 and 는지 for the reason that the ending is not read after them
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddChains(t *testing.T) {
	ix := &index{}
	ix.addChains()
	if len(ix.chainAux) != len(closingChains) || ix.graphs != len(closingChains) || ix.edges != len(closingChains) {
		t.Fatal(len(ix.chainAux), ix.graphs, ix.edges)
	}
	e := ix.chainAux["맺음:거나"]
	if e == nil || e.name != chainAuxName+"맺음:거나" {
		t.Fatal("the graph of 거나", e)
	}
	for _, before := range []string{"맺음:다", "맺음:ㄴ다"} {
		if v, _ := verdictOf(e.g, rule.Piece{Form: "거나", Before: before, HasStem: true, Next: -1}, "", false); v <= 0 {
			t.Errorf("%s takes 거나: %v", before, v)
		}
	}
	for _, before := range []string{"맺음:고", "맺음:지", "맺음:는지"} {
		if v, why := verdictOf(e.g, rule.Piece{Form: "거나", Before: before, HasStem: true, Next: -1}, "", true); v > 0 || !slices.Contains(why, whyNotAfter) {
			t.Errorf("%s takes 거나: %v %v", before, v, why)
		}
	}
}
