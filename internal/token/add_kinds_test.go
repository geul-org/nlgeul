//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddKinds — a graph for each kind, each with its exception, two rebuttals, the exception of a spoken form against the first and the exception of a sole word against the second, on the verb's and the adjective's the exceptions of a modifying closing and of a determiner before the stem with what lifts them, and on the graph of words that only stand alone the same two exceptions (48 edges), and on all of them the rebuttal that the stem is a misspelt form of another word, which lowers its verdict and keeps it above 0: right after a modifying closing or a determiner those three graphs turn a stem down and the others judge it as before, and the verb's and the adjective's do not turn down a listed auxiliary (boundAux) after a modifying closing
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddKinds(t *testing.T) {
	ix := &index{}
	ix.addKinds()
	if ix.graphs != 6 || ix.edges != 48 || len(ix.kinds) != 6 || ix.seen == nil {
		t.Fatal(ix.graphs, ix.edges)
	}
	fits := map[int]int{kindNominal: rule.TailCase, kindVerb: rule.TailPred, kindAdj: rule.TailAdj, kindAdverb: rule.TailAux, kindDeterminer: rule.TailBare, kindAlone: rule.TailBare}
	for _, kind := range kindOrder {
		free, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: "가르", Tail: fits[kind], Count: seenEnough}, "", false)
		after, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: "가르", Tail: fits[kind], Count: seenEnough, Adn: true}, "", false)
		det, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: "가르", Tail: fits[kind], Count: seenEnough, Det: true}, "", false)
		barred := kind&(kindVerb|kindAdj|kindAlone) != 0
		if free <= 0 || (barred && (after > 0 || det > 0)) || (!barred && (after != free || det != free)) {
			t.Fatal(kindLabel[kind], free, after, det)
		}
	}
	for _, kind := range kindOrder {
		whole := rule.Piece{Form: "가르", Tail: fits[kind], Count: seenEnough}
		typed := whole
		typed.Misspelt = true
		free, _ := verdictOf(ix.kinds[kind], whole, "", false)
		lowered, _ := verdictOf(ix.kinds[kind], typed, "", false)
		if lowered <= 0 || lowered >= free {
			t.Fatal(kindLabel[kind], "a misspelt stem stands lower, above 0", free, lowered)
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
