//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestRomanLetters — a text of the Roman numeral letters of Unicode alone (Ⅱ, ⅳ, ⅩⅢ) is one; a text with a Latin capital, a digit or another letter, and an empty text, are not
package parse

import "testing"

func TestRomanLetters(t *testing.T) {
	for text, want := range map[string]bool{"Ⅱ": true, "ⅳ": true, "ⅩⅢ": true, "Ⅰ": true, "ⅿ": true, "II": false, "ⅡA": false, "Ⅱ2": false, "①": false, "": false} {
		if got := romanLetters(text); got != want {
			t.Errorf("romanLetters(%q) = %v", text, got)
		}
	}
}
