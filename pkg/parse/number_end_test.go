//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNumberEnd — letters that end in a digit, or in greater numbers right after a digit (17억, 3천만), end in a number; greater numbers with no digit before them, other letters and nothing do not
package parse

import "testing"

func TestNumberEnd(t *testing.T) {
	for in, want := range map[string]bool{"17": true, "17억": true, "3천만": true, "억": false, "곰억": false, "17곰": false, "": false} {
		if got := numberEnd([]rune(in)); got != want {
			t.Errorf("%q: %v, want %v", in, got, want)
		}
	}
}
