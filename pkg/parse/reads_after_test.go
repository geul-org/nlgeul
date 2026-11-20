//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReadsAfter — the pieces the tables read after a number, by place: 만 is a greater number and no unit, 조 both, 년 and 도 units; a case mark, a particle and no letters are none
package parse

import "testing"

func TestReadsAfter(t *testing.T) {
	sc := scanOf("")
	cases := []struct {
		form  string
		place int
		reads bool
	}{
		{"만", placeBig, true},
		{"만", placeUnit, false},
		{"만", 0, true},
		{"조", placeBig, true},
		{"조", placeUnit, true},
		{"년", 0, true},
		{"도", placeUnit, true},
		{"도", placeBig, false},
		{"회합", 0, false},
		{"의", 0, false},
		{"", 0, false},
	}
	for _, c := range cases {
		if got := sc.readsAfter(c.form, c.place); got != c.reads {
			t.Errorf("%s at %d: %v, want %v", c.form, c.place, got, c.reads)
		}
	}
}
