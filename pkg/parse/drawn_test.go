//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestDrawn — a stretch of one sign is drawn with when its run of signs holds a sign that is no punctuation mark of a sentence: an arrow's head, a table's bar, another box-drawing letter, a sign further along the run (<!--); not next to a full stop, a question mark, a bracket or a quotation mark, not next to letters, and not by its own sign
package parse

import "testing"

func TestDrawn(t *testing.T) {
	for _, c := range []struct {
		text        string
		start, stop int
		want        bool
	}{
		{"-->", 0, 2, true},
		{"<--", 1, 3, true},
		{"<!--", 2, 4, true},
		{"|---|", 1, 4, true},
		{"├──┼", 1, 3, true},
		{"┌─┐", 1, 2, true},
		{"가--=나", 1, 3, true},
		{"다---.", 1, 4, false},
		{"─.”", 0, 1, false},
		{")——?", 1, 3, false},
		{"가---나", 1, 4, false},
		{"---", 0, 3, false},
		{"---", 1, 2, false},
		{"가─", 1, 2, false},
	} {
		if got := drawn([]rune(c.text), c.start, c.stop); got != c.want {
			t.Errorf("%q %d-%d: %v", c.text, c.start, c.stop, got)
		}
	}
}
