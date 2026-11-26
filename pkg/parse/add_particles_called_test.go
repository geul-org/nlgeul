//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestAddParticlesCalled — the vocative 야 is given the graph the word before it must pass, and the two candidates of the shortened particle 서 the graph that keeps it off a stem it makes a word of the word table with (joinedGate); the suffix 이 read where one is called the graph its name must pass (nameGate); no other particle has a gate, and the three graphs are counted once each
package parse

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
			called := e.name == vocName && e.gate == ix.call
			joined := e.form == "서" && e.gate.name == joinedGate && e.gate.g != nil
			named := e.name == nameSuffix && e.called && e.gate.name == nameGate && e.gate.g != nil
			if !called && !joined && !named {
				t.Errorf("%s has a gate", e.name)
			}
		}
	}
	if gated != 4 || ix.graphs != entries+3 {
		t.Fatal(gated, ix.graphs, entries)
	}
}
