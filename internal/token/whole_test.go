//ff:func feature=token type=reader control=sequence
//ff:what TestWhole — a word of no markers is read at the unit's end from every place, and inside the unit only when it is a known word that may stand alone or the script changes there, and never inside an atom
package token

import "testing"

func TestWhole(t *testing.T) {
	sc := scanOf("것곰")
	best := [][]*path{{{verdict: 1}}, nil, nil}
	if got := sc.whole(nil, best, 1); len(got) != 1 || got[0].sig != "것[의존명사:것]" {
		t.Fatal(got)
	}
	best[1] = sc.whole(nil, best, 1)
	if got := sc.whole(nil, best, 2); len(got) != 2 {
		t.Fatal(got)
	}
	if got := scanOf("곰것").whole(nil, best[:2], 1); len(got) != 0 {
		t.Fatal(got)
	}
	mixed := scanOf("API호출")
	if got := mixed.whole(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil, nil}, 3); len(got) != 1 || got[0].sig != "API[열:영숫자]" {
		t.Fatal("no word where the script changes", got)
	}
	if got := mixed.whole(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil, nil}, 2); len(got) != 0 {
		t.Fatal("a word ending inside an atom", got)
	}
	ws := wordsScan("사과먹", layerKind)
	if got := ws.whole(nil, [][]*path{{{verdict: 1}}, nil, nil}, 2); len(got) != 1 || got[0].sig != "사과[열쇠:체언]" {
		t.Fatal(got)
	}
	if got := wordsScan("먹사과", layerKind).whole(nil, [][]*path{{{verdict: 1}}, nil}, 1); len(got) != 0 {
		t.Fatal("a predicate stem standing alone", got)
	}
}
