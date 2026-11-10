//ff:func feature=token type=reader control=sequence
//ff:what TestSlotsOf — the different kinds once each, in order
package token

import "testing"

func TestSlotsOf(t *testing.T) {
	got := slotsOf([]*entry{{slot: 3}, {slot: 1}, {slot: 3}})
	if len(got) != 2 || got[0] != 1 || got[1] != 3 || slotsOf(nil) != nil {
		t.Fatal(got)
	}
}
