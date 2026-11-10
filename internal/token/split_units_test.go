//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSplitUnits — a run of digits and a Latin unit is split in two; a unit alone, digits alone and another run stay
package token

import "testing"

func TestSplitUnits(t *testing.T) {
	s := []rune("20ms kg 30 v2")
	got := splitUnits(s, runsOf(s))
	want := [][2]int{{0, 2}, {2, 4}, {5, 7}, {8, 10}, {11, 13}}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatal(got)
		}
	}
}
