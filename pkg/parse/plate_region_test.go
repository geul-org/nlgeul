//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPlateRegion — the name of a region of the table at a place of the unit is told by its length (서울, 경기 at the head, 부산 past a sign), and none where other letters stand, where the unit ends before the name does, or past the unit's end
package parse

import "testing"

func TestPlateRegion(t *testing.T) {
	cases := []struct {
		text string
		from int
		want int
	}{
		{"서울34가5678", 0, 2},
		{"경기1나2345", 0, 2},
		{"(부산7다8901", 1, 2},
		{"나라34가5678", 0, 0},
		{"서", 0, 0},
		{"서울", 2, 0},
		{"34가5678", 0, 0},
	}
	for _, c := range cases {
		sc := &scan{ix: scanOf("").ix, s: []rune(c.text)}
		if got := sc.plateRegion(c.from); got != c.want {
			t.Errorf("%s from %d: %d, want %d", c.text, c.from, got, c.want)
		}
	}
}
