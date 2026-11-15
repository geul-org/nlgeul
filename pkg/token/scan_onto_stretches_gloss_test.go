//ff:func feature=token type=reader control=iteration dimension=3
//ff:what TestOntoStretchesGloss — the stretches the net reads for a Hangul word written onto a quotation mark at a unit's head and glossed right after with as many Han characters ('관치 before (官治), as 그가 '대곤' '관치(官治)'라 cuts it) have no marker (the scan of the letters after the sign is told so: no 관하+치) and keep the word as a new word; told no gloss, they have 관하+치
package token

import (
	"strings"
	"testing"
)

func TestOntoStretchesGloss(t *testing.T) {
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
		var shown []string
		for _, st := range unitScan(r.ix, u).ontoStretches(u) {
			for _, n := range st.cells.nodes {
				for _, b := range n.branches {
					shown = append(shown, b.String())
				}
			}
		}
		all := "\n" + strings.Join(shown, "\n") + "\n"
		if strings.Contains(all, "관하[열쇠:동사]+치[맺음:지]") != c.marker || !strings.Contains(all, "\n관치\n") {
			t.Errorf("gloss %d: marker %v\n%s", c.gloss, c.marker, all)
		}
	}
}
