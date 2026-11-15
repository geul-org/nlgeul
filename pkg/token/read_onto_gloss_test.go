//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReadOntoGloss — a unit with a quotation mark at its head that may close and a Hangul word written onto it, glossed right after with as many Han characters ('관치 before (官治), as 그가 '대곤' '관치(官治)'라 cuts it), reads no marker in the letters after the sign either (the scan of those letters is told so: no 관하+치) and keeps the word as a new word and as known words; told no gloss, the same unit reads 관하+치
package token

import (
	"strings"
	"testing"
)

func TestReadOntoGloss(t *testing.T) {
	r := streamReader(t, "words-gloss.tsv")
	u := unitsOf(r.ix, "그가 '대곤' '관치(官治)'라 했다.", false)[4]
	if u.text != "'관치" || u.head != 1 || u.gloss != 2 {
		t.Fatal("the unit", u)
	}
	for _, c := range []struct {
		gloss  int
		marker bool
	}{{2, false}, {0, true}} {
		u.gloss = c.gloss
		got, _ := readOnto(r.ix, u, false)
		var lines []string
		for _, ct := range got {
			lines = append(lines, render(ct, 0)...)
		}
		all := "\n" + strings.Join(lines, "\n") + "\n"
		if strings.Contains(all, "관하[열쇠:동사]+치[맺음:지]") != c.marker || !strings.Contains(all, "\n  관치\t") || !strings.Contains(all, "\n  관[열쇠:체언] 치[열쇠:체언]\t") {
			t.Errorf("gloss %d: marker %v\n%s", c.gloss, c.marker, all)
		}
	}
}
