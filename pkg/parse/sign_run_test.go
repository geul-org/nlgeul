//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestSignRun — a word of one stem with no name whose letters are all signs is a run of signs (○○, □□□, ♠♠♠, ☞☞, **); a stem with a name, a letter or a digit among the signs, a marker after the run, an empty stem, a part that is no stem and no part at all are not
package parse

import "testing"

func TestSignRun(t *testing.T) {
	stem := func(text string) Part { return Part{Text: text, Stem: true} }
	for _, c := range []struct {
		parts []Part
		want  bool
	}{
		{[]Part{stem("○○")}, true},
		{[]Part{stem("□□□")}, true},
		{[]Part{stem("♠♠♠")}, true},
		{[]Part{stem("☞☞")}, true},
		{[]Part{stem("**")}, true},
		{[]Part{{Text: "○○", Stem: true, Name: Name{Kind: "열쇠", Form: "체언"}}}, false},
		{[]Part{stem("○가")}, false},
		{[]Part{stem("○1")}, false},
		{[]Part{stem("○○"), {Text: "가", Name: Name{Kind: "주격", Form: "가"}}}, false},
		{[]Part{stem("")}, false},
		{[]Part{{Text: "○○"}}, false},
		{nil, false},
	} {
		if got := signRun(c.parts); got != c.want {
			t.Errorf("%+v: %v, want %v", c.parts, got, c.want)
		}
	}
}
