//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSurvivors — the states that stay when the net is kept from a place on: a state that begins there stays when a root is before it, a state that begins after it when a state that stays is before it, the others (before the place, with no root or state that stays before them) do not
package token

import "testing"

func TestSurvivors(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	order := r.byEnd()
	roots := map[*path]bool{}
	for _, s := range order {
		if s.last.end == 4 && s.start == 0 {
			roots[s] = true
			break
		}
	}
	alive := survivors(order, 4, roots)
	if len(alive) == 0 {
		t.Fatal("none stays")
	}
	for s := range alive {
		if s.start < 4 {
			t.Fatal("begins before", s.start)
		}
	}
	if got := survivors(order, 4, map[*path]bool{}); len(got) != 0 {
		t.Fatal("no root", len(got))
	}
}
