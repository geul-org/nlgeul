//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitNames — the name column of the unit table gives every unit its own letters for now: each name is the letters of the unit, and every unit has one
package token

import "testing"

func TestUnitNames(t *testing.T) {
	if len(unitWords) == 0 || len(unitNames) != len(unitWords) {
		t.Fatal(len(unitWords), len(unitNames))
	}
	for _, u := range unitWords {
		if name, ok := unitNames[u]; !ok || name != u {
			t.Errorf("%s: name %q", u, name)
		}
	}
}
