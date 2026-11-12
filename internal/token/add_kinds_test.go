//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddKinds — a graph for each kind, each with its exception and two rebuttals, and on the verb's and the adjective's the exception of a modifying closing before the stem with the two against it (24 edges): right after a modifying closing those two graphs turn a stem down and the others judge it as before, and neither turns down a listed auxiliary (boundAux)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddKinds(t *testing.T) {
	ix := &index{}
	ix.addKinds()
	if ix.graphs != 6 || ix.edges != 24 || len(ix.kinds) != 6 || ix.seen == nil {
		t.Fatal(ix.graphs, ix.edges)
	}
	fits := map[int]int{kindNominal: rule.TailCase, kindVerb: rule.TailPred, kindAdj: rule.TailAdj, kindAdverb: rule.TailAux, kindDeterminer: rule.TailBare, kindAlone: rule.TailBare}
	for _, kind := range kindOrder {
		free, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: "가르", Tail: fits[kind], Count: seenEnough}, "", false)
		after, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: "가르", Tail: fits[kind], Count: seenEnough, Adn: true}, "", false)
		pred := kind&(kindVerb|kindAdj) != 0
		if free <= 0 || (pred && after > 0) || (!pred && after != free) {
			t.Fatal(kindLabel[kind], free, after)
		}
	}
	for _, form := range boundAux {
		verb, _ := verdictOf(ix.kinds[kindVerb], rule.Piece{Form: form, Tail: rule.TailPred, Count: seenEnough, Adn: true}, "", false)
		adj, _ := verdictOf(ix.kinds[kindAdj], rule.Piece{Form: form, Tail: rule.TailPred, Count: seenEnough, Adn: true}, "", false)
		if verb <= 0 || adj <= 0 {
			t.Fatal(form, verb, adj)
		}
	}
}
