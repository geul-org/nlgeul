//ff:func feature=token type=reader control=sequence
//ff:what TestNetRunTryAt — a stretch from the last cut to a place is cut when its best way is the spans wanted, and not when it is another way
package token

import "testing"

func TestNetRunTryAt(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	if _, ok := r.tryAt(4, []Span{{0, 2}}); ok || r.c != 0 || r.stats.cuts != 0 {
		t.Fatal("another way")
	}
	st, ok := r.tryAt(4, []Span{{0, 4}})
	if !ok || r.c != 4 || st.contest.Text != "고양이가" || st.cells == nil {
		t.Fatal(ok, r.c)
	}
}
