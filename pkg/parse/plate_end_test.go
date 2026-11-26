//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPlateEnd — the shape of a registration number: a kind of two or three digits, a letter of the table and a serial of four digits ends after the serial (also before the unit 호 and particles, after the name of a region with a kind of one or two digits, and from a place past the unit's head); none for a kind of other lengths (one digit or four with no region, three after a region), a letter the table lacks, a serial of other lengths, a serial that a Latin letter runs on from, a greater number (60조5000), and a unit with another piece read after a number right after the serial (100주5000원), but the unit 호 after the serial of a letter that is a unit is the plate's (12호3456호)
package parse

import "testing"

func TestPlateEnd(t *testing.T) {
	cases := []struct {
		text string
		from int
		end  int
	}{
		{"306너8412", 0, 8},
		{"74모2238", 0, 7},
		{"74모2238호를", 0, 7},
		{"서울34가5678", 0, 9},
		{"경기1나2345", 0, 8},
		{"(306너8412)", 1, 9},
		{"12호3456호", 0, 7},
		{"12호3456", 0, 7},
		{"3너8412", 0, 0},
		{"3061너8412", 0, 0},
		{"서울345가5678", 0, 0},
		{"306힣8412", 0, 0},
		{"306너841", 0, 0},
		{"306너84125", 0, 0},
		{"306너8412a", 0, 0},
		{"306너", 0, 0},
		{"60조5000", 0, 0},
		{"100주5000원", 0, 0},
		{"너8412", 0, 0},
	}
	for _, c := range cases {
		sc := &scan{ix: scanOf("").ix, s: []rune(c.text)}
		if got := sc.plateEnd(c.from); got != c.end {
			t.Errorf("%s from %d: %d, want %d", c.text, c.from, got, c.end)
		}
	}
}
