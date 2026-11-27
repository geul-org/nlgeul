//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNamedSign — a sign of the sign table (a comma, ㆍ) and a sign the tables of special letters name by its shape (㈎, ⓐ, ？) have a name of their own; a sign in no table (○), a letter of the tables that is no sign (①) and a Hangul syllable have none
package parse

import "testing"

func TestNamedSign(t *testing.T) {
	for r, want := range map[rune]bool{',': true, 'ㆍ': true, '㈎': true, 'ⓐ': true, '？': true, '○': false, '①': false, '가': false} {
		if got := namedSign(r); got != want {
			t.Errorf("namedSign(%c) = %v", r, got)
		}
	}
}
