//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestScanBetween — the states from one place to another: the states that end at the last place and every state before them, the root or a seed is none, a state before several states once, in the order they were made
package token

import "testing"

func TestScanBetween(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	pt := r.sc.between(0, 8)
	if pt.from != 0 || pt.to != 8 || len(pt.states) == 0 {
		t.Fatal(pt)
	}
	seen := map[*path]bool{}
	ends := 0
	for _, s := range pt.states {
		if seen[s] || s.last.end > 8 || s.prev == nil {
			t.Fatal("state", s.last.sig)
		}
		seen[s] = true
		if s.last.end == 8 {
			ends++
		}
	}
	if ends == 0 {
		t.Fatal("no state ends at 8")
	}
}
