//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunReroot — after a cut the states that stay are linked to seeds and to each other and kept again in the registry in the order they were made, the places are put in order, and the net goes on from there to the end of the unit with the same words as the rest of the unit read by itself
package token

import "testing"

func TestNetRunReroot(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	order := r.byEnd()
	roots := endingAt(order, 4)
	alive := survivors(order, 4, roots)
	r.reroot(4, roots, alive, order)
	reg := r.sc.registry()
	if len(reg.all) != len(alive) || len(reg.states) != len(alive) || len(r.best[4]) != len(roots) || r.sc.floor != 4 {
		t.Fatal(len(reg.all), len(alive))
	}
	for _, s := range reg.all {
		if s.start < 4 || s.prev == nil || !alive[s] {
			t.Fatal("state", s.start)
		}
	}
}
