//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestGlossedBy — letters that are Hangul syllables alone, as many as the Han characters glossed after them, with no sign trimmed between, are glossed (대곤 by 2, 운용한다 by 4); no gloss, a sign between, a count that differs, and a letter that is no Hangul syllable (a digit, a Latin letter, a lone jamo, a Han character) are not
package token

import "testing"

func TestGlossedBy(t *testing.T) {
	for _, c := range []struct {
		letters string
		han     int
		after   string
		want    bool
	}{
		{"대곤", 2, "", true},
		{"관치", 2, "", true},
		{"운용한다", 4, "", true},
		{"대곤", 0, "", false},
		{"", 0, "", false},
		{"대곤", 2, ",", false},
		{"대곤", 1, "", false},
		{"대곤", 3, "", false},
		{"3곤", 2, "", false},
		{"A곤", 2, "", false},
		{"ㄱ곤", 2, "", false},
		{"大곤", 2, "", false},
	} {
		if got := glossedBy([]rune(c.letters), c.han, c.after); got != c.want {
			t.Errorf("glossedBy(%q, %d, %q) = %v, want %v", c.letters, c.han, c.after, got, c.want)
		}
	}
}
