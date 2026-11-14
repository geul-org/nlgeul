//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunAgreed — nothing is agreed on while a word that may begin at the last cut is still open, later the first spans every state at the window's end begins with, of which every way goes on after the last
package token

import "testing"

func TestNetRunAgreed(t *testing.T) {
	r, _ := windowSteps(t, 5, 6)
	if got := r.agreed(); got != nil {
		t.Fatal("open", got)
	}
	r, _ = windowSteps(t, 5, 24)
	got := r.agreed()
	if len(got) == 0 || got[0] != (Span{0, 4}) {
		t.Fatal(got)
	}
	front, _ := r.front()
	for _, ln := range r.lines(front) {
		if len(ln) <= len(got) {
			t.Fatal("a way ends with the part", ln, got)
		}
	}
}
