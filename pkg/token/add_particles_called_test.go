//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestAddParticlesCalled — the vocative 야 alone is given the graph the word before it must pass, and that graph is counted once (a graph and an edge)
package token

import "testing"

func TestAddParticlesCalled(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addParticles()
	if ix.call == nil || ix.call.name != callGate {
		t.Fatal(ix.call)
	}
	gated, entries := 0, 0
	for _, list := range ix.byKey {
		for _, e := range list {
			entries++
			if e.gate == nil {
				continue
			}
			gated++
			if e.name != vocName || e.gate != ix.call {
				t.Errorf("%s has a gate", e.name)
			}
		}
	}
	if gated != 1 || ix.graphs != entries+1 {
		t.Fatal(gated, ix.graphs, entries)
	}
}
