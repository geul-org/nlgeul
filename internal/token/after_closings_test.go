//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAfterClosings — 도 right after the linking closing 지 is one more piece of the predicate's tail (also after a dropped copula); not after a closing that takes none (다), not for a pair that is a closing of the tables itself (어도), not for a particle that is not one of the few, nor inside a longer chain; 어서 and 면서 take one too (먹+어서+는, 가+면서+는) but 으면서도, a closing itself, is read as that closing only; with a word table any particle, inside a chain too, right after the naming closing 기 is one more piece under its own name (먹+기+는, 먹+기+에+는), and without one 기 takes none; end to end a noun of the table that ends in 기 still stands as the noun, and a stem no table has takes no 기
package token

import (
	"strings"
	"testing"
)

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
	neun0 := piece{text: "는", names: "보조사:는", verdict: 1}
	if got = scanOf("먹어서는").afterClosings(3, 0, neun0, []*entry{{name: "보조사:는"}}); len(got) == 0 || got[0].pieces[0].names != "맺음:어서" || got[0].pieces[1].names != "보조사:는" {
		t.Fatal("어서 takes an auxiliary particle", got)
	}
	if got = scanOf("가면서는").afterClosings(3, 0, neun0, []*entry{{name: "보조사:는"}}); len(got) == 0 || got[0].pieces[0].names != "맺음:면서" {
		t.Fatal("면서 takes an auxiliary particle", got)
	}
	if got = scanOf("먹으면서도").afterClosings(4, 0, do, aux); len(got) != 0 {
		t.Fatal("으면서도 is a closing of the tables", got)
	}
	if got = wordsScan("보지도", layerCost).afterClosings(2, 0, do, aux); len(got) == 0 || got[0].pieces[0].names != "맺음:지" || got[0].pieces[1].names != "보조사:도" {
		t.Fatal("with a word table 지 still takes 도", got)
	}

	eu := piece{text: "에", names: "부사격:에", verdict: 1}
	if got = scanOf("먹기에").afterClosings(2, 1, eu, []*entry{{name: "부사격:에"}}); got != nil {
		t.Fatal("without a word table 기 takes no particle", got)
	}
	naming := 0
	for _, c := range wordsScan("먹기에", layerCost).afterClosings(2, 1, eu, []*entry{{name: "부사격:에"}}) {
		if c.start == 1 && len(c.pieces) == 2 && c.pieces[0].names == "이름:기" && c.pieces[1] == eu && c.gate != nil {
			naming++
		}
	}
	if naming != 1 {
		t.Fatal("기 and 에 after it, inside a chain, asking for a predicate's stem", naming)
	}
	neun := piece{text: "는", names: "보조사:는", verdict: 1}
	if got = wordsScan("먹기는", layerCost).afterClosings(2, 0, neun, []*entry{{name: "보조사:는"}}); len(got) == 0 || got[0].pieces[0].names != "이름:기" || got[0].pieces[1] != neun {
		t.Fatal("기 and 는 after it", got)
	}

	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["고기"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	for in, want := range map[string]string{
		"먹기는":  "먹[열쇠:동사]+기[이름:기]+는[보조사:는]",
		"먹기에는": "먹[열쇠:동사]+기[이름:기]+에[부사격:에]+는[보조사:는]",
		"고기는":  "고기[열쇠:체언]+는[보조사:는]",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	if got := strings.Join(r.Lines("곰기는", false), "\n"); strings.Contains(got, "[이름:기]") {
		t.Fatal("기 after a stem no table has\n" + got)
	}
	if got := strings.Join(Lines("먹기는 먹기에는", false), "\n"); strings.Contains(got, "[이름:기]") {
		t.Fatal("기 without a word table\n" + got)
	}
}
