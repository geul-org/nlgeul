//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLineBare — an input line holds no letter and no digit when it is signs and spaces alone (a rule, a table's line, rows of dashes), told of the line the place stands in and of no other line of the sentence
package parse

import "testing"

func TestLineBare(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want bool
	}{
		{"---", 0, true},
		{"|---|---|", 3, true},
		{"---- ----", 5, true},
		{"가 --- 나", 2, false},
		{"--- 3", 0, false},
		{"제목\n---\n글", 3, true},
		{"제목\n---\n글", 5, true},
		{"제목\n---\n글", 0, false},
		{"제목\n---\n글", 7, false},
		{"---\n글", 1, true},
	} {
		if got := newCutter(ix, c.text, false).lineBare(c.at); got != c.want {
			t.Errorf("%q at %d: %v", c.text, c.at, got)
		}
	}
}
