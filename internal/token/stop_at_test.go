//ff:func feature=token type=reader control=sequence
//ff:what TestStopAt — a full stop after a word closes it; dots in a row are read once, at their end, as 줄임표; a dot inside a name, at the unit's start, or no dot is no stop
package token

import "testing"

func TestStopAt(t *testing.T) {
	sc := scanOf("했다.그리고")
	if p, start, ok := sc.stopAt(3); !ok || start != 2 || p.names != "마침표:." || p.text != "." {
		t.Fatal(p, start, ok)
	}
	sc = scanOf("글쎄...모르")
	if p, start, ok := sc.stopAt(5); !ok || start != 2 || p.names != "줄임표:." || p.text != "..." {
		t.Fatal(p, start, ok)
	}
	if _, _, ok := sc.stopAt(4); ok {
		t.Fatal("inside a row of dots")
	}
	if _, _, ok := scanOf("task.py").stopAt(5); ok {
		t.Fatal("a dot inside a name")
	}
	if _, _, ok := scanOf(".가").stopAt(1); ok {
		t.Fatal("at the unit's start")
	}
	if _, _, ok := sc.stopAt(1); ok {
		t.Fatal("no dot")
	}
	if _, _, ok := sc.stopAt(0); ok {
		t.Fatal("place 0")
	}
}
