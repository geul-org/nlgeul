//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestAmountEnd — letters end in an amount when they end in a number (a digit, greater numbers written onto digits) or in a word of the unit table right after one; not in a word, not in a unit with no number before it, not when empty
package parse

import "testing"

func TestAmountEnd(t *testing.T) {
	for text, want := range map[string]bool{"100": true, "17억": true, "3,000원": true, "50%": true, "3천만원": true, "원": false, "가격": false, "가격원": false, "": false} {
		if got := amountEnd([]rune(text)); got != want {
			t.Errorf("%q: %v", text, got)
		}
	}
}
