//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunFront — the states at the window's end are those that end where a word may still begin that the cut must not split; none while a word may begin at the last cut, and none when no state ends there
package token

import "testing"

func TestNetRunFront(t *testing.T) {
	r, _ := windowSteps(t, 5, 6)
	if got, ok := r.front(); ok || got != nil {
		t.Fatal("open at the cut")
	}
	r, _ = windowSteps(t, 5, 24)
	got, ok := r.front()
	if !ok || len(got) == 0 {
		t.Fatal(got)
	}
	for _, s := range got {
		if r.sc.tokens(s.last.end, r.j) >= windowSize {
			t.Fatal("too far", s.last.end)
		}
	}
	r.sc.registry().all = nil
	if got, ok = r.front(); ok || got != nil {
		t.Fatal("no state")
	}
}
