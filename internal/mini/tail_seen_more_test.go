//ff:func feature=mini type=rule control=sequence
//ff:what TestTailSeenMore — the greatest share among the tails as long as the chain; shorter tails and unknown counts give no value
package mini

import "testing"

func TestTailSeenMore(t *testing.T) {
	w := &word{chain: chainFact{size: 2, count: 1}, tails: []tailFact{{size: 2, count: 1}, {size: 2, count: 3}, {size: 1, count: 99}}}
	v, _ := tailSeenMore(ctxFor(w, tailFact{}), nil)
	w.tails = []tailFact{{size: 1, count: 99}, {size: 2}}
	none, _ := tailSeenMore(ctxFor(w, tailFact{}), nil)
	if v != 0.5 || none != -1 {
		t.Fatal(v, none)
	}
}
