//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCuts — a plus between letters cuts the sentence; not a plus at a unit's edge or next to a space, a run of them, one next to another sign, one in a span of code, nor one inside an address, a path or an option; the trace tells the graphs' judgments
package token

import "testing"

func TestCuts(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want bool
	}{
		{"곰+돌", 1, true},
		{"4+5", 1, true},
		{"곰+8", 1, true},
		{"+3", 0, false},
		{"가+", 1, false},
		{"가 +나", 2, false},
		{"가+ 나", 1, false},
		{"가++나", 1, false},
		{"가++나", 2, false},
		{"b+=7", 1, false},
		{"C++", 1, false},
		{"가+(나", 1, false},
		{"`a+b`", 2, false},
		{"/a/b+c/d", 4, false},
		{"http://a.com/x+y", 14, false},
		{"--a+b", 3, true},
		{"가-나", 1, false},
	} {
		c0 := newCutter(ix, c.text, false)
		if got, notes := c0.cuts(c.at); got != c.want || len(notes) != 0 {
			t.Errorf("%q at %d: %v %v, want %v", c.text, c.at, got, notes, c.want)
		}
	}
	traced := newCutter(ix, "곰+돌", true)
	if got, notes := traced.cuts(1); !got || len(notes) != 2 || notes[0] != "  · +\t더하기\t1.00\t통과" || notes[1] != "  · +\t연결\t1.00\t통과" {
		t.Fatal("the trace", got, notes)
	}
	barred := newCutter(ix, "가++나", true)
	if _, notes := barred.cuts(1); len(notes) != 2 || notes[0] != "  · +\t더하기\t-1.00\t기각: 같은 기호가 이어진다, 다른 기호와 맞닿아 있다" {
		t.Fatal("a run is judged and falls", notes)
	}
	next := newCutter(ix, "b+=7", true)
	if _, notes := next.cuts(1); len(notes) != 2 || notes[0] != "  · +\t더하기\t-1.00\t기각: 다른 기호와 맞닿아 있다" {
		t.Fatal("next to another sign", notes)
	}
}
