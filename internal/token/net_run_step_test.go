//ff:func feature=token type=reader control=sequence
//ff:what TestNetRunStep — a place is worked out (the states made so far are noted as the most at once), the words are planted while no cut has been made and no more after it, and the most states is only ever raised
package token

import "testing"

func TestNetRunStep(t *testing.T) {
	r, _ := windowSteps(t, 1, 5)
	if r.j != 5 || r.stats.states == 0 || r.stats.states != len(r.sc.registry().all) {
		t.Fatal(r.j, r.stats.states)
	}
	r.stats.states = 1 << 30
	r.step()
	if r.stats.states != 1<<30 {
		t.Fatal("only raised")
	}
	planted := 0
	r.plant = func(j int) []*path { planted++; return nil }
	r.step()
	r.c = 1
	r.step()
	if planted != 1 || r.j != 8 {
		t.Fatal(planted, r.j)
	}
}
