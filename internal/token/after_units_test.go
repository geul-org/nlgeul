//ff:func feature=token type=reader control=sequence
//ff:what TestAfterUnits — a tail is read after the unit that ends where it starts and keeps its closing; a tail fused into a stem, and one with no unit before it, give nothing
package token

import "testing"

func TestAfterUnits(t *testing.T) {
	sc := scanOf("3개다")
	e := &entry{name: "맺음:다"}
	got := sc.afterUnits(tail{start: 2, pieces: []piece{{text: "다"}}, close: e, noYo: true})
	if len(got) != 1 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].close != e || !got[0].noYo {
		t.Fatal(got)
	}
	if len(sc.afterUnits(tail{start: 2, base: '가'})) != 0 || len(sc.afterUnits(tail{start: 1})) != 0 {
		t.Fatal("read where it should not be")
	}
}
