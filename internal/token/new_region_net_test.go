//ff:func feature=token type=helper control=sequence
//ff:what TestNewRegionNet — a region from cells of a stretch that begins at 2 of the unit and lies at 10 of the sentence: its span, its nodes at the spans moved by 8 with their letters, Glued (true for a node that does not begin the stretch, else the stretch's), the candidates as given, the words that follow each, the best and the rest, the number of ways, the clashes, a member test that takes spans of the sentence, and the frontier with the words that begin it
package token

import (
	"math/big"
	"testing"
)

func TestNewRegionNet(t *testing.T) {
	br := func(text string, p float64) []Branch {
		return []Branch{{Words: []Eojeol{{Parts: []Part{{Text: text, Stem: true}}}}, Verdict: 1, Pref: p, scores: []wordScore{{1, p, 0}}}}
	}
	cs := cells{
		from:    2,
		nodes:   []cell{{Span{2, 3}, br("가", -1)}, {Span{2, 4}, br("가나", -3)}, {Span{3, 5}, br("나다", -2)}, {Span{4, 5}, br("다", -1)}},
		edges:   []edge{{Span{2, 3}, Span{3, 5}}, {Span{2, 4}, Span{4, 5}}},
		ways:    big.NewInt(2),
		member:  func(spans []Span) bool { return lineKey(spans) == "2-3,3-5" },
		clashes: 1,
	}
	rg := newRegionNet(Contest{Unit: 4, Glued: true, Text: "가나다"}, 10, cs)
	if rg.span != (Span{10, 13}) || len(rg.nodes) != 4 || rg.nodes[0].Span != (Span{10, 11}) || rg.nodes[2].Span != (Span{11, 13}) || rg.ways.Int64() != 2 || rg.clashes != 1 {
		t.Fatal(rg.span, rg.nodes)
	}
	if c := rg.nodes[2].Contest; c.Text != "나다" || c.Unit != 4 || !c.Glued || rg.nodes[0].Contest.Text != "가" || !rg.nodes[0].Contest.Glued {
		t.Fatal("letters and Glued", rg.nodes)
	}
	if free := newRegionNet(Contest{Text: "가나다"}, 10, cs); free.nodes[0].Contest.Glued || !free.nodes[2].Contest.Glued {
		t.Fatal("Glued is the stretch's for the node that begins it")
	}
	if len(rg.next[0]) != 1 || rg.next[0][0] != 2 || len(rg.next[1]) != 1 || rg.next[1][0] != 3 || len(rg.front.items) != 2 {
		t.Fatal(rg.next, rg.front.items)
	}
	if rg.best[0] != -1 || rg.rest[0] != -3 || rg.rest[1] != -4 || rg.rest[2] != -2 {
		t.Fatal(rg.best, rg.rest)
	}
	if !rg.member([]Span{{10, 11}, {11, 13}}) || rg.member([]Span{{10, 12}, {12, 13}}) {
		t.Fatal("member is asked with spans of the sentence")
	}
}
