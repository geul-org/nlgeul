//ff:func feature=token type=reader control=sequence
//ff:what TestNewNetRun — a net run of a unit has its places ready (the root at place 0, the length, the window on, the numbers of birth from 0); the index made to work without the window gives a run with no window; a scan with no letters has nothing to work out
package token

import "testing"

func TestNewNetRun(t *testing.T) {
	sc, u, r0 := netUnit(t, "words-far.tsv", windowText(1))
	stats := &netStats{}
	r := newNetRun(sc, u, true, nil, stats)
	if r.n != len(sc.s) || r.done || !r.window || len(r.best) != r.n+1 || len(r.best[0]) != 1 || r.base != 0 || !r.first || r.stats != stats {
		t.Fatal(r)
	}
	r0.ix.noWindow = true
	if newNetRun(unitScan(r0.ix, u), u, true, nil, stats).window {
		t.Fatal("the window is off")
	}
	empty := newNetRun(&scan{ix: r0.ix}, unit{}, true, nil, stats)
	if !empty.done || empty.best != nil {
		t.Fatal("nothing to work out")
	}
}
