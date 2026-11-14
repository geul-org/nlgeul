//ff:func feature=token type=reader control=sequence
//ff:what TestNetEnd — the net of a unit closed from its first place is the net of the whole unit: its states that end the unit, from place 0
package token

import "testing"

func TestNetEnd(t *testing.T) {
	sc, _, _ := netUnit(t, "words-far.tsv", "고양이가마신다")
	lat := sc.net()
	if lat == nil || lat.from != 0 || lat.n != 7 || len(lat.ends) == 0 {
		t.Fatal(lat)
	}
}
