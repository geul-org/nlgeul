//ff:func feature=parse type=reader control=iteration dimension=3
//ff:what TestOntoReaderGloss — a unit with a quotation mark at its head that may close and a Hangul word written onto it, glossed right after with as many Han characters ('관치 before (官治), as 그가 '대곤' '관치(官治)'라 cuts it), has no marker in the stretches of the letters after the sign (the scan of those letters is told so: no 관하+치) and keeps the word as a new word; told no gloss, the stretches have 관하+치
package parse

import (
	"strings"
	"testing"
)

func TestOntoReaderGloss(t *testing.T) {
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
		rd := ontoReader(r.ix, u, false, &netStats{})
		var shown []string
		for st, ok := rd.next(); ok; st, ok = rd.next() {
			if st.cells == nil {
				continue
			}
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
