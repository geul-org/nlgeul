//ff:func feature=token type=reader control=sequence
//ff:what TestBeginsOf — the places a word with its stem end at e may begin at are those worked out (beginnings) when the scan has them, and every place before e when it does not; the units of the tests are read by the net too
package token

import (
	"slices"
	"testing"
)

func TestBeginsOf(t *testing.T) {
	driveAll(t)
	sc := scanOf("가나다라")
	sc.beginnings()
	if got := sc.beginsOf(3); !slices.Equal(got, sc.begins[3]) || !slices.Equal(got, []int{0, 1, 2}) {
		t.Fatal("places worked out", got)
	}
	sc.begins = nil
	if got := sc.beginsOf(3); !slices.Equal(got, []int{0, 1, 2}) || len(sc.beginsOf(0)) != 0 {
		t.Fatal("places not worked out", got)
	}
}
