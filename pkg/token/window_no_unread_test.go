//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWindowNoUnread — a window that would cut where every state at its end ends is not cut there but widened: after a made-up run of letters nothing reads (nine letters of no word, then a particle) in a long unit, the stream with the window reads to the end with no stretch that nothing reads, none blocked, and the best way is that of the stream without the window
package token

import (
	"strings"
	"testing"
)

func TestWindowNoUnread(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	a := "아버지가방에들어가신다"
	text := strings.Repeat(a, 3) + "뷁뷁뷁뷁뷁뷁뷁뷁뷁를" + strings.Repeat(a, 3)
	r.ix.noWindow = false
	st := r.Stream(text)
	for sg, ok := st.Next(); ok; sg, ok = st.Next() {
		if len(sg.Nodes) == 0 {
			t.Fatal("a stretch nothing reads", sg.Region)
		}
	}
	if st.stats.blocked != 0 || st.stats.unread != 0 {
		t.Fatal(st.stats)
	}
	on, off := flatNodes(r, text, true), flatNodes(r, text, false)
	if len(on) == 0 || strings.Join(on, "\n") != strings.Join(off, "\n") {
		t.Fatal("\n", on, "\n", off)
	}
}
