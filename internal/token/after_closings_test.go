//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAfterClosings — 도 right after the linking closing 지 is one more piece of the predicate's tail (also after a dropped copula); not after a closing that takes none (다), not for a pair that is a closing of the tables itself (어도), not for a particle that is not one of the few, nor inside a longer chain; 어서 and 면서 take one too (먹+어서+는, 가+면서+는) but 으면서도, a closing itself, is read as that closing only; with a word table any particle, inside a chain too, right after the naming closing 기 is one more piece under its own name (먹+기+는, 먹+기+에+는), and without one 기 takes none; before a particle after 기 a dropped copula stands before the 기 (문제+(이)+기+도), the tails are looked up there as chained, and the same lookup not chained has no such copula; a case particle right after a closing that asks is one more piece under its own name, inside a chain too (먹+는지+를, 먹+는지+에 before 는), also after a closing under a final consonant (가+(ㄹ)지+도) and after the copula (학교+이+(ㄴ)지+가), while 의, which such a closing does not take, and 를 after the plain 지 are not read so; end to end a noun of the table that ends in 기 still stands as the noun, and a stem no table has takes no 기
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestAfterClosings(t *testing.T) {
	do := piece{text: "도", names: "보조사:도", verdict: 1}
	aux := &entry{name: "보조사:도"}
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
	if got = scanOf("보지도").afterClosings(2, 0, do, &entry{name: "주격:가"}); got != nil {
		t.Fatal("not an auxiliary particle of the few", got)
	}
	if got = scanOf("보지도").afterClosings(2, 1, do, aux); got != nil {
		t.Fatal("inside a longer chain", got)
	}
	neun0 := piece{text: "는", names: "보조사:는", verdict: 1}
	if got = scanOf("먹어서는").afterClosings(3, 0, neun0, &entry{name: "보조사:는"}); len(got) == 0 || got[0].pieces[0].names != "맺음:어서" || got[0].pieces[1].names != "보조사:는" {
		t.Fatal("어서 takes an auxiliary particle", got)
	}
	if got = scanOf("가면서는").afterClosings(3, 0, neun0, &entry{name: "보조사:는"}); len(got) == 0 || got[0].pieces[0].names != "맺음:면서" {
		t.Fatal("면서 takes an auxiliary particle", got)
	}
	if got = scanOf("먹으면서도").afterClosings(4, 0, do, aux); len(got) != 0 {
		t.Fatal("으면서도 is a closing of the tables", got)
	}
	if got = wordsScan("보지도", layerCost).afterClosings(2, 0, do, aux); len(got) == 0 || got[0].pieces[0].names != "맺음:지" || got[0].pieces[1].names != "보조사:도" {
		t.Fatal("with a word table 지 still takes 도", got)
	}

	eu := piece{text: "에", names: "부사격:에", verdict: 1}
	if got = scanOf("먹기에").afterClosings(2, 1, eu, &entry{name: "부사격:에"}); got != nil {
		t.Fatal("without a word table 기 takes no particle", got)
	}
	naming := 0
	for _, c := range wordsScan("먹기에", layerCost).afterClosings(2, 1, eu, &entry{name: "부사격:에"}) {
		if c.start == 1 && len(c.pieces) == 2 && c.pieces[0].names == "이름:기" && c.pieces[1] == eu && c.gate != nil {
			naming++
		}
	}
	if naming != 1 {
		t.Fatal("기 and 에 after it, inside a chain, asking for a predicate's stem", naming)
	}
	neun := piece{text: "는", names: "보조사:는", verdict: 1}
	if got = wordsScan("먹기는", layerCost).afterClosings(2, 0, neun, &entry{name: "보조사:는"}); len(got) == 0 || got[0].pieces[0].names != "이름:기" || got[0].pieces[1] != neun {
		t.Fatal("기 and 는 after it", got)
	}

	ws := wordsScan("문제기도", layerCost)
	dropped := 0
	for _, c := range ws.afterClosings(3, 0, do, aux) {
		if c.start == 2 && len(c.pieces) == 3 && c.pieces[0].text == "(이)" && c.pieces[0].names == droppedName && c.pieces[1].names == "이름:기" && c.pieces[2] == do {
			dropped++
		}
	}
	if dropped != 1 {
		t.Fatal("a dropped copula before 기 and 도 after it", dropped)
	}
	for _, chained := range []bool{false, true} {
		got := ws.closings(3, false, chained)
		if has := slices.ContainsFunc(got, func(c tail) bool { return c.pieces[0].names == droppedName }); has != chained {
			t.Fatal("a dropped copula before 기 only in the lookup that is chained", chained, got)
		}
	}

	reul := piece{text: "를", names: "목적격:를", verdict: 1}
	obj := &entry{name: "목적격:를"}
	if got = scanOf("먹는지를").afterClosings(3, 0, reul, obj); len(got) == 0 || got[0].start != 1 || got[0].pieces[0].names != "맺음:는지" || got[0].pieces[1] != reul {
		t.Fatal("는지 takes 를", got)
	}
	if got = scanOf("먹는지에").afterClosings(3, 1, eu, &entry{name: "부사격:에"}); len(got) == 0 || got[0].pieces[0].names != "맺음:는지" || got[0].pieces[1] != eu {
		t.Fatal("는지 takes 에 inside a chain", got)
	}
	if got = scanOf("먹는지의").afterClosings(3, 0, piece{text: "의", names: "관형격:의", verdict: 1}, &entry{name: "관형격:의"}); len(got) != 0 {
		t.Fatal("는지 takes no 의", got)
	}
	if got = scanOf("먹지를").afterClosings(2, 0, reul, obj); len(got) != 0 {
		t.Fatal("the plain 지 takes no case particle", got)
	}
	if got = scanOf("먹었다를").afterClosings(3, 0, reul, obj); got != nil {
		t.Fatal("no closing that asks ends before 를", got)
	}

	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["고기"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	for in, want := range map[string]string{
		"먹기는":    "먹[열쇠:동사]+기[이름:기]+는[보조사:는]",
		"먹기에는":   "먹[열쇠:동사]+기[이름:기]+에[부사격:에]+는[보조사:는]",
		"고기는":    "고기[열쇠:체언]+는[보조사:는]",
		"먹는지를":   "먹[열쇠:동사]+는지[맺음:는지]+를[목적격:를]",
		"작은지는":   "작[열쇠:형용사]+은지[맺음:은지]+는[보조사:는]",
		"갈지도":    "가[열쇠:동사]+(ㄹ)지[맺음:ㄹ지]+도[보조사:도]",
		"먹을지가":   "먹[열쇠:동사]+을지[맺음:을지]+가[주격:가]",
		"학교인지와":  "학교[열쇠:체언]+이[서술격:이]+(ㄴ)지[맺음:ㄴ지]+와[",
		"먹었는지에는": "먹[열쇠:동사]+었[선어말:었]+는지[맺음:는지]+에[부사격:에]+는[보조사:는]",
	} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	if got := strings.Join(r.Lines("곰기는", false), "\n"); strings.Contains(got, "[이름:기]") {
		t.Fatal("기 after a stem no table has\n" + got)
	}
	for _, in := range []string{"먹지를", "먹는지의", "먹는다를"} {
		if got := strings.Join(r.Lines(in, false), "\n"); strings.Contains(got, "]+를[목적격:를]") && strings.Contains(got, "[맺음:지]+를") || strings.Contains(got, "[맺음:는지]+의") || strings.Contains(got, "다]+를") {
			t.Fatal("a particle after a closing that takes none\n" + got)
		}
	}
	bare, _ := New(Options{})
	if got := strings.Join(bare.Lines("먹기는 먹기에는", false), "\n"); strings.Contains(got, "[이름:기]") {
		t.Fatal("기 without a word table\n" + got)
	}
	if got := strings.Join(bare.Lines("먹는지를", false), "\n"); !strings.Contains(got, "먹+는지[맺음:는지]+를[목적격:를]") {
		t.Fatal("는지 and 를 without a word table\n" + got)
	}
}
