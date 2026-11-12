//ff:func feature=token type=reader control=sequence
//ff:what TestWhole — a new word after a boundary does not stand where its letters read as markers (the scan's marks at j), the unit's first word does; a word of no markers is read at the unit's end from every place, and inside the unit only when it is a known word that may stand alone or the script changes there, and never inside an atom; nor where its letters read as two runs of markers one after the other (는 then 게), while one run is the old rule and letters with no run stand; end to end 곰이는게 has no new word 는게 after 곰+이
package token

import (
	"strings"
	"testing"
)

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
	after := scanOf("곰에서")
	kept := [][]*path{{{verdict: 1}}, {{verdict: 1}}, nil, nil}
	after.marks[3] = map[int]bool{1: true}
	if got := after.whole(nil, kept, 3); len(got) != 1 || got[0].sig != "곰에서" {
		t.Fatal("a new word after a boundary whose letters read as markers", got)
	}
	after.marks[3] = map[int]bool{0: true}
	if got := after.whole(nil, kept, 3); len(got) != 2 {
		t.Fatal("the unit's first word is not asked", got)
	}
	runs := scanOf("곰는게")
	start := [][]*path{{{verdict: 1}}, {{verdict: 1}}, nil, nil}
	if got := runs.whole(nil, start, 3); len(got) != 2 {
		t.Fatal("letters with no run stand", got)
	}
	runs.marks[3] = map[int]bool{1: true}
	if got := runs.whole(nil, start, 3); len(got) != 1 || got[0].sig != "곰는게" {
		t.Fatal("one run, by the old rule", got)
	}
	runs.marks[2] = map[int]bool{1: true}
	runs.marks[3] = map[int]bool{2: true}
	if got := runs.whole(nil, start, 3); len(got) != 1 || got[0].sig != "곰는게" {
		t.Fatal("는 and then 게: two runs", got)
	}
	r, _ := New(Options{})
	got := strings.Join(r.Lines("곰이는게", false), "\n")
	if strings.Contains(got, "곰+이[주격:이|보격:이] 는게\t") || !strings.Contains(got, "곰이는게\t") {
		t.Fatal("곰이는게\n" + got)
	}
}
