//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSoundAt — letters alone in a square or round bracket right after Han characters are told how many characters stand there (江山[강산]: 2, 明月(명월): 2, a Latin letter before the characters stops the count: K君[군]: 1); no Han character right before the bracket, a space between, a bracket that is not closed by its own sign right after the letters, another kind of bracket, the letters' place at the sentence's head and an end past the sentence are told 0
package token

import "testing"

func TestSoundAt(t *testing.T) {
	for _, c := range []struct {
		text    string
		at, end int
		want    int
	}{
		{"江山[강산]에", 3, 5, 2},
		{"明月(명월)이", 3, 5, 2},
		{"K君[군]", 3, 4, 1},
		{"그 山[산]", 4, 5, 1},
		{"강산[강산]", 3, 5, 0},
		{"江山 [강산]", 4, 6, 0},
		{"江山[강산)", 3, 5, 0},
		{"江山(강산]", 3, 5, 0},
		{"江山{강산}", 3, 5, 0},
		{"江山[강산]", 3, 4, 0},
		{"[강]", 1, 2, 0},
		{"江山[강산", 3, 5, 0},
	} {
		if got := soundAt([]rune(c.text), c.at, c.end); got != c.want {
			t.Errorf("soundAt(%q, %d, %d) = %d, want %d", c.text, c.at, c.end, got, c.want)
		}
	}
}
