//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSignRune — punctuation marks and symbols are signs; letters, digits and spaces are not
package token

import "testing"

func TestSignRune(t *testing.T) {
	for r, want := range map[rune]bool{'+': true, '·': true, ',': true, '$': true, '가': false, 'a': false, '3': false, ' ': false} {
		if signRune(r) != want {
			t.Errorf("%q: want %v", r, want)
		}
	}
}
