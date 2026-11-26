//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCaseNumberEnd — the shape of a court case number: a year of two or four digits, a mark of one to three syllables and a serial ends after the serial (also before a unit and particles: 2008회합9호로, and from a place past the unit's head); none for a year of other lengths, no mark, a mark of four syllables, no serial, a serial that Latin letters run on from, a greater number, a mark that ends with a joint between two numbers (11의2, 31조의4, 3000분의7, 36조제1항), and a unit with another piece after a number right after the serial (2020년3월); a unit that is a case mark too is a mark with nothing after the serial (36도5, 2019차3141), any other unit is a unit (11초92, 10시30, 81승53)
package parse

import "testing"

func TestCaseNumberEnd(t *testing.T) {
	cases := []struct {
		text string
		from int
		end  int
	}{
		{"2004다25123", 0, 10},
		{"97도2836", 0, 7},
		{"94누8853은", 0, 7},
		{"2008회합9호로", 0, 7},
		{"2010카기전12", 0, 9},
		{")2004다1", 1, 7},
		{"1다2", 0, 0},
		{"123다4", 0, 0},
		{"2004다", 0, 0},
		{"2004가나다라1", 0, 0},
		{"2004다25123a", 0, 0},
		{"다25123", 0, 0},
		{"17만2000", 0, 0},
		{"60조5000억", 0, 0},
		{"11의2", 0, 0},
		{"31조의4", 0, 0},
		{"3000분의7", 0, 0},
		{"36조제1항", 0, 0},
		{"2020년3월", 0, 0},
		{"36도5분", 0, 0},
		{"36도5", 0, 4},
		{"2019차3141", 0, 9},
		{"11초92", 0, 0},
		{"10시30", 0, 0},
		{"81승53", 0, 0},
	}
	for _, c := range cases {
		sc := &scan{ix: scanOf("").ix, s: []rune(c.text)}
		if got := sc.caseNumberEnd(c.from); got != c.end {
			t.Errorf("%s from %d: %d, want %d", c.text, c.from, got, c.end)
		}
	}
}
