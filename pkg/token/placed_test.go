//ff:func feature=token type=helper control=sequence
//ff:what TestPlaced — anything stands with nothing to its right; places go up; greater numbers follow one another, other places do not repeat
package token

import "testing"

func TestPlaced(t *testing.T) {
	if !placed(placeUnit, 0) || !placed(placeUnit, placeOrder) || placed(placeOrder, placeUnit) || !placed(placeBig, placeBig) || placed(placeUnit, placeUnit) {
		t.Fatal("placed")
	}
}
