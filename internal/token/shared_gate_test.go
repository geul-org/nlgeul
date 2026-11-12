//ff:func feature=token type=helper control=sequence
//ff:what TestSharedGate — the gate of some readings is the one they all ask for: that of a lone reading, that of readings with the same gate, and none when one reading asks for none or for another
package token

import "testing"

func TestSharedGate(t *testing.T) {
	a, b := &entry{name: "a"}, &entry{name: "b"}
	if sharedGate([]*entry{{gate: a}}) != a || sharedGate([]*entry{{}}) != nil {
		t.Fatal("a lone reading's gate is its own")
	}
	if sharedGate([]*entry{{gate: a}, {gate: a}}) != a {
		t.Fatal("readings with one gate share it")
	}
	if sharedGate([]*entry{{gate: a}, {}}) != nil || sharedGate([]*entry{{}, {gate: a}}) != nil {
		t.Fatal("a reading that asks for none leaves the form with none")
	}
	if sharedGate([]*entry{{gate: a}, {gate: b}}) != nil {
		t.Fatal("readings with different gates share none")
	}
}
