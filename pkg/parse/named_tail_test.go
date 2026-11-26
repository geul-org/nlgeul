//ff:func feature=parse type=reader control=sequence
//ff:what TestNamedTail — the whole unit of a path with Hangul after its last slash is a path by that reading, lowered when the Hangul reads as markers alone; the letters of a path with no head and one slash are a path up to where they end; a stem that begins elsewhere or ends elsewhere, and a unit with neither reading, are none
package parse

import "testing"

func TestNamedTail(t *testing.T) {
	sc := scanOf("~/a/b/회의록")
	sc.search()
	if got := sc.namedTail(0, 9); len(got) != 1 || got[0].name != pathName || got[0].v != sc.ix.pathNamed {
		t.Fatal(got)
	}
	if got := append(sc.namedTail(0, 6), sc.namedTail(1, 9)...); got != nil {
		t.Fatal(got)
	}
	sc = scanOf("~/a/b/는")
	sc.search()
	if got := sc.namedTail(0, 7); len(got) != 1 || got[0].v != sc.ix.pathMarked {
		t.Fatal(got)
	}
	sc = scanOf("notes/회의-v2를")
	sc.search()
	if got := sc.namedTail(0, 11); len(got) != 1 || got[0].name != pathName || got[0].v != sc.ix.pathWeak {
		t.Fatal(got)
	}
	if got := append(sc.namedTail(0, 12), sc.namedTail(6, 11)...); got != nil {
		t.Fatal(got)
	}
	sc = scanOf("고양이가")
	sc.search()
	if got := sc.namedTail(0, 4); got != nil {
		t.Fatal(got)
	}
}
