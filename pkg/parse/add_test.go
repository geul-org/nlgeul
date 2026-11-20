//ff:func feature=parse type=reader control=sequence
//ff:what TestAdd — an entry goes under its key with a graph, and the totals grow; an entry whose name the case exception table has gets the exception of its rows (one edge more), which turns down the word of a row
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdd(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.add(&entry{name: "부사격:에서", form: "에서", class: classParticle, count: 5})
	if len(ix.byKey['서']) != 1 || ix.byKey['서'][0].g == nil || ix.graphs != 1 || ix.edges != 3 || ix.total != 5 {
		t.Fatalf("%+v", ix)
	}
	useCases(t, "testdata/case-exceptions.tsv")
	plain, held := &entry{name: "맺음:다", form: "다", class: classClosing}, &entry{name: "맺음:소", form: "소", class: classClosing}
	ix.add(plain)
	edges := ix.edges
	ix.add(held)
	if ix.edges != edges+3 || len(ix.byKey['소']) != 1 {
		t.Fatal(ix.edges-edges, ix.byKey['소'])
	}
	p := rule.Piece{Form: "소", HasStem: true, Prev: '먹', Next: -1}
	if v, _ := verdictOf(held.g, p, "먹소", false); v != -1 {
		t.Fatal("the word of a row", v)
	}
	if v, _ := verdictOf(held.g, p, "적소", false); v != 1 {
		t.Fatal("another word", v)
	}
}
