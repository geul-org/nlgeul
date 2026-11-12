//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosingDotsHead — after a head (a closing bracket, a closing quotation mark, a span of code) a sign dropped between the head and the core does not hide the dots that close the unit: )-나. ends in a full stop, "-이다... in an ellipsis, `a`-다. in a full stop, as without the sign; a unit with a head that does not end in dots has none; end to end (가)-나. and "인용"-이다... end in those lines
package token

import "testing"

func TestClosingDotsHead(t *testing.T) {
	r, _ := New(Options{})
	for _, c := range []struct {
		field string
		head  int
		dots  string
		want  string
	}{
		{")-나.", 1, ".", ".[마침표:.]"},
		{")나.", 1, ".", ".[마침표:.]"},
		{"\"-이다...", 1, "...", "...[줄임표:.]"},
		{"`a`-다.", 3, ".", ".[마침표:.]"},
	} {
		sc := &scan{ix: r.ix, s: coreOf(unit{text: c.field, head: c.head}), head: c.head}
		got := sc.closingDots(c.field)
		if len(got) != 1 || got[0].Text != c.dots || render(got[0])[0] != c.want || got[0].Branches[0].Words[0].Parts[0].Stem {
			t.Errorf("%s: %+v", c.field, got)
		}
	}
	if got := (&scan{ix: r.ix, s: coreOf(unit{text: ")-나", head: 1}), head: 1}).closingDots(")-나"); got != nil {
		t.Fatal("a unit with a head that does not end in dots", got)
	}
	for in, want := range map[string]string{"(가)-나.": ".[마침표:.]", "\"인용\"-이다...": "...[줄임표:.]"} {
		if lines := r.Lines(in, false); len(lines) == 0 || lines[len(lines)-1] != want {
			t.Errorf("%s: the last line is not %s\n%v", in, want, lines)
		}
	}
}
