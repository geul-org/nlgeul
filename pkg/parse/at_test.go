//ff:func feature=parse type=reader control=sequence
//ff:what TestAt — a letter of the unit, and 0 outside it
package parse

import "testing"

func TestAt(t *testing.T) {
	sc := &scan{s: []rune("곰이")}
	if sc.at(1) != '이' || sc.at(-1) != 0 || sc.at(2) != 0 {
		t.Fatal("at")
	}
}
