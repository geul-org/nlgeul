//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestOntoStretches — the letters written onto a sign of a pair at a unit's head (a closing sign and a particle, an opening sign and a word with a particle) are read by the net: the stretches cover those letters and nothing of the sign, each with cells (read by the net, not as a contest), and they are the stretches readOntoNet puts after the sign's own contest
package token

import (
	"strings"
	"testing"
)

func TestOntoStretches(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, u := range []unit{
		{text: ")를", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}},
		{text: "(고양이가", head: 1, name: pairName, candidates: []pairing{{openName, 1}}},
	} {
		got := unitScan(ix, u).ontoStretches(u)
		var texts []string
		for _, st := range got {
			if st.cells == nil || st.redo != nil || len(st.contest.Branches) != 0 {
				t.Fatal(u.text, "not a stretch of the net", st.contest.Text)
			}
			texts = append(texts, st.contest.Text)
		}
		all, _ := readOntoNet(ix, u, false)
		if strings.Join(texts, "") != string([]rune(u.text)[u.head:]) || len(all) != len(got)+1 || all[1].contest.Text != got[0].contest.Text {
			t.Fatal(u.text, texts, len(all))
		}
	}
}
