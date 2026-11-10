//ff:func feature=token type=reader control=sequence
//ff:what TestWhole — a new word after a boundary does not stand where its letters read as markers, the unit's first word does; a word of no markers is read at the unit's end from every place, and inside the unit only when it is a known word that may stand alone or the script changes there, and never inside an atom
package token

import "testing"

func TestWhole(t *testing.T) {
	sc := scanOf("것곰")
	best := [][]*path{{{verdict: 1}}, nil, nil}
	if got := sc.whole(nil, best, 1, nil); len(got) != 1 || got[0].sig != "것[의존명사:것]" {
		t.Fatal(got)
	}
	best[1] = sc.whole(nil, best, 1, nil)
	if got := sc.whole(nil, best, 2, nil); len(got) != 2 {
		t.Fatal(got)
	}
	if got := scanOf("곰것").whole(nil, best[:2], 1, nil); len(got) != 0 {
		t.Fatal(got)
	}
	mixed := scanOf("API호출")
	if got := mixed.whole(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil, nil}, 3, nil); len(got) != 1 || got[0].sig != "API[열:영숫자]" {
		t.Fatal("no word where the script changes", got)
	}
	if got := mixed.whole(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil, nil}, 2, nil); len(got) != 0 {
		t.Fatal("a word ending inside an atom", got)
	}
	ws := wordsScan("사과먹", layerKind)
	if got := ws.whole(nil, [][]*path{{{verdict: 1}}, nil, nil}, 2, nil); len(got) != 1 || got[0].sig != "사과[열쇠:체언]" {
		t.Fatal(got)
	}
	if got := wordsScan("먹사과", layerKind).whole(nil, [][]*path{{{verdict: 1}}, nil}, 1, nil); len(got) != 0 {
		t.Fatal("a predicate stem standing alone", got)
	}
	after := scanOf("곰에서")
	kept := [][]*path{{{verdict: 1}}, {{verdict: 1}}, nil, nil}
	if got := after.whole(nil, kept, 3, map[int]bool{1: true}); len(got) != 1 || got[0].sig != "곰에서" {
		t.Fatal("a new word after a boundary whose letters read as markers", got)
	}
	if got := after.whole(nil, kept, 3, map[int]bool{0: true}); len(got) != 2 {
		t.Fatal("the unit's first word is not asked", got)
	}
}
