//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestBracketedAt — letters stand alone in a round bracket written right onto a word when the bracket opens right before them, right after a Hangul syllable, a Latin letter or a digit, and closes right after them; not after a space or at the sentence's head, in another bracket, or with more letters before the closing sign
package parse

import "testing"

func TestBracketedAt(t *testing.T) {
	for _, c := range []struct {
		text   string
		i, end int
		want   bool
	}{
		{"크로뮴(III)", 4, 7, true},
		{"Fe(II)", 3, 5, true},
		{"2(II)", 2, 4, true},
		{"크로뮴 (III)", 5, 8, false},
		{"(III)", 1, 4, false},
		{"크로뮴[III]", 4, 7, false},
		{"크로뮴(III", 4, 7, false},
		{"크로뮴(III 가)", 4, 7, false},
		{"크로뮴,(III)", 5, 8, false},
	} {
		if got := bracketedAt([]rune(c.text), c.i, c.end); got != c.want {
			t.Errorf("bracketedAt(%q, %d, %d) = %v", c.text, c.i, c.end, got)
		}
	}
}
