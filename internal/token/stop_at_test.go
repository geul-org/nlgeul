//ff:func feature=token type=reader control=sequence
//ff:what TestStopAt — a full stop after a word closes it; dots in a row are read once, at their end, as 줄임표; a comma closes the word before it like a dot; a dot inside a name, a dot or a comma at the unit's start, commas in a row, or no dot is no stop
package token

import "testing"

func TestStopAt(t *testing.T) {
	sc := scanOf("했다.그리고")
	if p, start := sc.stopAt(3); len(p) != 1 || start != 2 || p[0].names != "마침표" || p[0].text != "." {
		t.Fatal(p, start)
	}
	sc = scanOf("글쎄...모르")
	if p, start := sc.stopAt(5); len(p) != 1 || start != 2 || p[0].names != "줄임표" || p[0].text != "..." {
		t.Fatal(p, start)
	}
	if p, _ := sc.stopAt(4); len(p) > 0 {
		t.Fatal("inside a row of dots")
	}
	if p, _ := scanOf("task.py").stopAt(5); len(p) > 0 {
		t.Fatal("a dot inside a name")
	}
	if p, _ := scanOf(".가").stopAt(1); len(p) > 0 {
		t.Fatal("at the unit's start")
	}
	if p, start := scanOf("사과,포도").stopAt(3); len(p) != 1 || start != 2 || p[0].names != "쉼표" || p[0].text != "," {
		t.Fatal("a comma closes the word before it", p, start)
	}
	if p, _ := scanOf(",가").stopAt(1); len(p) > 0 {
		t.Fatal("a comma at the unit's start")
	}
	if p, _ := scanOf("가,,나").stopAt(3); len(p) > 0 {
		t.Fatal("commas in a row")
	}
	if p, _ := sc.stopAt(1); len(p) > 0 {
		t.Fatal("no dot")
	}
	if p, _ := sc.stopAt(0); len(p) > 0 {
		t.Fatal("place 0")
	}
}
