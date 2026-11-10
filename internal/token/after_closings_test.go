//ff:func feature=token type=reader control=sequence
//ff:what TestAfterClosings — 도 right after the linking closing 지 is one more piece of the predicate's tail (also after a dropped copula); not after a closing that takes none (다), not for a pair that is a closing of the tables itself (어도), not for a particle that is not one of the few, nor inside a longer chain
package token

import "testing"

func TestAfterClosings(t *testing.T) {
	do := piece{text: "도", names: "보조사:도", verdict: 1}
	aux := []*entry{{name: "보조사:도"}}
	got := scanOf("보지도").afterClosings(2, 0, do, aux)
	if len(got) != 2 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].pieces[0].names != "맺음:지" || got[0].pieces[1].names != "보조사:도" || got[0].close == nil {
		t.Fatal(got)
	}
	if got = scanOf("보았다도").afterClosings(3, 0, do, aux); len(got) != 0 {
		t.Fatal("after a closing that takes no particle", got)
	}
	if got = scanOf("먹어도").afterClosings(2, 0, do, aux); len(got) != 0 {
		t.Fatal("a pair that is a closing of the tables", got)
	}
	if got = scanOf("보지도").afterClosings(2, 0, do, []*entry{{name: "주격:가"}}); got != nil {
		t.Fatal("not an auxiliary particle of the few", got)
	}
	if got = scanOf("보지도").afterClosings(2, 1, do, aux); got != nil {
		t.Fatal("inside a longer chain", got)
	}
}
