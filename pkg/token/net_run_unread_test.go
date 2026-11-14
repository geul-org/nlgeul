//ff:func feature=token type=reader control=sequence
//ff:what TestNetRunUnread — the letters between two places are a stretch with no reading: no node, no way, the letters as they are, and it is counted
package token

import "testing"

func TestNetRunUnread(t *testing.T) {
	r, _ := windowRun(t, "words-far.tsv", windowText(1))
	st := r.unread(4, 8)
	rg := newRegionNet(Contest{Text: st.contest.Text}, 4, *st.cells)
	if st.contest.Text != "마당에서" || r.stats.unread != 1 || st.cells.ways.Sign() != 0 || len(rg.nodes) != 0 {
		t.Fatal(st.contest.Text, r.stats.unread)
	}
	if _, ok := rg.deal(); ok {
		t.Fatal("no way")
	}
}
