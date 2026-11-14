//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNumeralBefore — the letters right before a place read as a numeral (삼십, 세); the head of a unit does not, nor letters that are none, nor a long run of letters of which none is a numeral
package token

import "testing"

func TestNumeralBefore(t *testing.T) {
	for _, c := range []struct {
		text string
		at   int
		want bool
	}{
		{"삼십개", 2, true},
		{"삼십개", 0, false},
		{"세마리", 1, true},
		{"의자개", 2, false},
		{"일가나다라마바거너더러머버개", 13, false},
		{"열두개", 2, true},
		{"이백삼십사만오천육백칠십팔원", 13, true},
	} {
		if got := scanOf(c.text).numeralBefore(c.at); got != c.want {
			t.Fatal(c.text, c.at, got)
		}
	}
}
