//ff:func feature=token type=helper control=sequence
//ff:what TestUnitStretchRegion — a stretch with cells becomes the region its cells make (a node at each cell); one without becomes the region of the ways its branches cut it in; both at the place of the sentence given
package token

import (
	"math/big"
	"testing"
)

func TestUnitStretchRegion(t *testing.T) {
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "가나", Stem: true}}, End: 2}}, Verdict: 1, Pref: -1}
	old := unitStretch{contest: Contest{Text: "가나", Branches: []Branch{br}}}.region(Contest{Text: "가나", Branches: []Branch{br}}, 5)
	if old.span != (Span{5, 7}) || len(old.nodes) != 1 || old.ways.Int64() != 1 || old.nodes[0].Span != (Span{5, 7}) {
		t.Fatal(old)
	}
	cs := &cells{from: 0, nodes: []cell{{Span{0, 2}, []Branch{br}}}, ways: big.NewInt(1), member: func([]Span) bool { return true }}
	net := unitStretch{contest: Contest{Text: "가나"}, cells: cs}.region(Contest{Text: "가나"}, 5)
	if net.span != (Span{5, 7}) || len(net.nodes) != 1 || net.nodes[0].Span != (Span{5, 7}) {
		t.Fatal(net)
	}
}
