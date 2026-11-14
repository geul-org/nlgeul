//ff:func feature=token type=reader control=sequence
//ff:what TestNetRunPrune — the net is kept from the cut on standing on the states given; when none stays (blocked) every state that ends at the cut is a root (widened), and when still none no stretch is said to be unread (finish does that at the unit's end)
package token

import "testing"

func TestNetRunPrune(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	r.prune(4, r.best[4])
	if r.stats.blocked != 0 || len(r.seeds) != len(r.best[4]) || r.done {
		t.Fatal(r.stats)
	}
	r, _ = windowSteps(t, 5, 24)
	r.prune(4, nil)
	if r.stats.blocked != 1 || r.stats.widened != 1 || r.done || len(r.seeds) == 0 {
		t.Fatal(r.stats)
	}
	r, _ = windowSteps(t, 5, 24)
	r.prune(24, nil)
	if r.stats.blocked != 1 || r.stats.widened != 0 || r.done || len(r.rest) != 0 || r.stats.unread != 0 || len(r.seeds) != 0 {
		t.Fatal(r.stats, r.rest)
	}
}
