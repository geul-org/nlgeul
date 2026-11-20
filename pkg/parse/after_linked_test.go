//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAfterLinked — a copula tail is read once more after the linking 어서 (먹+어서, 가+(아)서) that ends where it starts: the closing comes first, then the copula's pieces, each now of a predicate tail, and the tail keeps its closing, 요 and its gate; nothing is read after another closing (먹고, 먹으면서), for a tail fused into a stem, at the head of the word, or where the letter before the tail ends no such closing
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAfterLinked(t *testing.T) {
	e := &entry{name: "맺음:다"}
	g := &entry{name: "앞말:체언"}
	pieces := []piece{{text: "이", names: copulaName, tail: rule.TailCopula}, {text: "다", names: "맺음:다", tail: rule.TailPred}}
	for _, text := range []string{"먹어서이다", "가서이다"} {
		start := len([]rune(text)) - 2
		got := scanOf(text).afterLinked(tail{start: start, pieces: pieces, close: e, noYo: true, gate: g})
		i := slices.IndexFunc(got, func(l tail) bool { return slices.Contains(copulaClosings, l.pieces[0].names) })
		if i < 0 {
			t.Fatalf("%s: no 어서 before the copula %v", text, got)
		}
		l := got[i]
		if len(l.pieces) != 3 || l.pieces[1].names != copulaName || l.pieces[2].names != "맺음:다" || l.close != e || !l.noYo || l.gate != g {
			t.Fatalf("%s: %+v", text, l)
		}
		if slices.ContainsFunc(l.pieces[1:], func(p piece) bool { return p.tail != rule.TailPred }) {
			t.Fatalf("%s: %+v are not all pieces of a predicate tail", text, l.pieces[1:])
		}
		if pieces[0].tail != rule.TailCopula {
			t.Fatal("the copula tail's own pieces are left as they were")
		}
	}
	for _, c := range []struct {
		text  string
		start int
		base  rune
	}{
		{"먹고이다", 2, 0}, {"먹어서이다", 3, '서'}, {"이다", 0, 0}, {"서이다", 1, 0}, {"먹다이다", 2, 0}, {"먹으면서이다", 4, 0},
	} {
		if got := scanOf(c.text).afterLinked(tail{start: c.start, pieces: pieces, base: c.base}); len(got) != 0 {
			t.Errorf("%s: %v", c.text, got)
		}
	}
}
