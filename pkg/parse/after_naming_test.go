//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAfterNaming — with a word table a copula tail is read once more after a naming closing that ends where it starts (먹+음+이+다, 모르+(ㅁ)+이+다): the closing comes first, then the copula's pieces, each now of a predicate tail, and the tail keeps its closing and 요; the gate a dropped stem asks of a nominal is not kept; nothing is read without a word table, for a tail fused into a stem, after a stem of one syllable (됨이다), where the letter before the tail ends no naming closing, after a closing that only ends in ㅁ (간담), or after another closing (먹고, and the naming 기: 먹기)
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAfterNaming(t *testing.T) {
	e := &entry{name: "맺음:다"}
	pieces := []piece{{text: "이", names: copulaName, tail: rule.TailCopula}, {text: "다", names: "맺음:다", tail: rule.TailPred}}
	for text, name := range map[string]string{"먹음이다": "이름:음", "모름이다": nounName} {
		sc := wordsScan(text, layerCost)
		start := len([]rune(text)) - 2
		got := sc.afterNaming(tail{start: start, pieces: pieces, close: e, noYo: true, gate: sc.ix.nominal})
		i := slices.IndexFunc(got, func(l tail) bool { return l.pieces[len(l.pieces)-3].names == name })
		if i < 0 {
			t.Fatalf("%s: no %s before the copula %v", text, name, got)
		}
		l := got[i]
		if l.pieces[len(l.pieces)-2].names != copulaName || l.pieces[len(l.pieces)-1].names != "맺음:다" || l.close != e || !l.noYo || l.gate == sc.ix.nominal {
			t.Fatalf("%s: %+v", text, l)
		}
		if slices.ContainsFunc(l.pieces, func(p piece) bool { return p.tail == rule.TailCopula }) || pieces[0].tail != rule.TailCopula {
			t.Fatalf("%s: %+v are not all pieces of a predicate tail, or the copula tail's own pieces changed", text, l.pieces)
		}
	}
	if got := scanOf("먹음이다").afterNaming(tail{start: 2, pieces: pieces, close: e}); len(got) != 0 {
		t.Fatal("without a word table", got)
	}
	if got := wordsScan("간담이다", layerCost).afterNaming(tail{start: 2, pieces: pieces}); slices.ContainsFunc(got, func(l tail) bool { return l.pieces[0].names == "맺음:ㄴ담" }) {
		t.Fatal("a closing that only ends in ㅁ is no naming closing", got)
	}
	for _, c := range []struct {
		text  string
		start int
		base  rune
	}{
		{"먹음이다", 2, '음'}, {"됨이다", 1, 0}, {"먹고이다", 2, 0}, {"먹다이다", 2, 0}, {"먹기이다", 2, 0},
	} {
		if got := wordsScan(c.text, layerCost).afterNaming(tail{start: c.start, pieces: pieces, base: c.base}); len(got) != 0 {
			t.Errorf("%s: %v", c.text, got)
		}
	}
}
