//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestEnglishNumber — a numeral of the table is a number written in English whatever the case of its letters (two, Two, TWO, hundred), and so are numerals joined by hyphens (twenty-one); another word, a numeral joined to another word, a part of a numeral and an empty text are not
package parse

import "testing"

func TestEnglishNumber(t *testing.T) {
	for text, want := range map[string]bool{"two": true, "Two": true, "TWO": true, "zero": true, "nineteen": true, "ninety": true, "hundred": true, "trillion": true, "twenty-one": true, "Forty-Two": true,
		"second": false, "tw": false, "twos": false, "one-on-one": false, "two-": false, "2": false, "": false} {
		if got := englishNumber(text); got != want {
			t.Errorf("englishNumber(%q) = %v", text, got)
		}
	}
}
