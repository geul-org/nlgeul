//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestRomanSure — a Roman numeral of I, V and X of two letters or more is not left open by its letters (II, IV, XIV, XXI); one letter (I, V, X) and X alone repeated (XX, XXX) are, and so is what is no such numeral
package parse

import "testing"

func TestRomanSure(t *testing.T) {
	for text, want := range map[string]bool{"II": true, "III": true, "IV": true, "XIV": true, "XXI": true, "XIX": true, "I": false, "V": false, "X": false, "XX": false, "XXX": false, "CD": false, "": false} {
		if got := romanSure(text); got != want {
			t.Errorf("romanSure(%q) = %v", text, got)
		}
	}
}
