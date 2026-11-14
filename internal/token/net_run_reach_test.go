//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunReach — every state kept has a best way from the states before the first: its value is its span's best lambda plus that of the state before it, the root or a seed is worth 0, and the states at a span are grouped
package token

import "testing"

func TestNetRunReach(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	rc := r.reach()
	all := r.sc.registry().all
	if len(rc.value) != len(all) || len(rc.back) != len(all) || len(rc.group) == 0 {
		t.Fatal(len(rc.value), len(all))
	}
	for _, s := range all {
		sp := Span{Start: s.start, End: s.last.end}
		prior := 0.0
		if b := rc.back[s]; b != nil && b.prev != nil {
			prior = rc.value[b]
		}
		if rc.value[s] != prior+r.lambdaAt(sp, rc.group[sp]) {
			t.Fatal("value", sp)
		}
	}
}
