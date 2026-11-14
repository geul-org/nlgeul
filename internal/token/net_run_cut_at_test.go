//ff:func feature=token type=reader control=sequence
//ff:what TestNetRunCutAt — a cut is made at the end of the first span of the agreed part that a stretch reads the same way; a part no stretch reads so gives none; where a sign shows a boundary at that end the candidates kept there are worked out first and the cut is tried again
package token

import "testing"

func TestNetRunCutAt(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	if _, ok := r.cutAt([]Span{{0, 2}}); ok || r.c != 0 {
		t.Fatal("no stretch reads so")
	}
	r.sc.mid[4] = midSign{}
	st, ok := r.cutAt(r.agreed())
	if !ok || st.contest.Text != "고양이가" || r.c != 4 {
		t.Fatal(st.contest.Text, r.c)
	}
	if _, done := r.sc.kept[4]; !done {
		t.Fatal("kept at the sign")
	}
}
