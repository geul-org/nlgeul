//ff:func feature=parse type=rule control=sequence
//ff:what TestAddNamed — each particle of closingParticles has a graph of its own, named with the particle's name, and the closings of its row by the same name; the count of graphs and edges grows by one each; the graph of the auxiliary 야 takes the closing 다 before it and turns down 고 for the reason that the particle is not read after it
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddNamed(t *testing.T) {
	ix := &index{}
	ix.addNamed()
	n := len(closingParticles)
	if n == 0 || len(ix.namedAux) != n || len(ix.namedEnds) != n || ix.graphs != n || ix.edges != n {
		t.Fatal(len(ix.namedAux), len(ix.namedEnds), ix.graphs, ix.edges)
	}
	e := ix.namedAux["보조사:야"]
	if e == nil || e.name != namedAuxName+"보조사:야" || !slices.Contains(ix.namedEnds["보조사:야"], "맺음:라") {
		t.Fatal("the graph of 야", e, ix.namedEnds["보조사:야"])
	}
	if v, _ := verdictOf(e.g, rule.Piece{Form: "야", Before: "맺음:다", HasStem: true, Next: -1}, "", false); v <= 0 {
		t.Fatal("다 takes 야", v)
	}
	if v, why := verdictOf(e.g, rule.Piece{Form: "야", Before: "맺음:고", HasStem: true, Next: -1}, "", true); v > 0 || !slices.Contains(why, whyNotAfter) {
		t.Fatal("고 takes 야", v, why)
	}
}
