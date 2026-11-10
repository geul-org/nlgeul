//ff:func feature=rule type=rule control=sequence
//ff:what TestTailSeenMore — the greatest share among the tails as long as the chain; shorter tails and unknown counts give no value
package rule

import "testing"

func TestTailSeenMore(t *testing.T) {
	w := &Word{Chain: Chain{Size: 2, Count: 1}, Tails: []Tail{{Size: 2, Count: 1}, {Size: 2, Count: 3}, {Size: 1, Count: 99}}}
	v, _ := TailSeenMore(ctxFor(w, Tail{}), nil)
	w.Tails = []Tail{{Size: 1, Count: 99}, {Size: 2}}
	none, _ := TailSeenMore(ctxFor(w, Tail{}), nil)
	if v != 0.5 || none != -1 {
		t.Fatal(v, none)
	}
}
