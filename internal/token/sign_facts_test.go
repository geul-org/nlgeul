//ff:func feature=token type=reader control=sequence
//ff:what TestSignFacts — the runs a sign joins are those tied by the same sign around it: in 2026-11-10~12 the hyphens see three runs, the range sign two
package token

import "testing"

func TestSignFacts(t *testing.T) {
	labels, signs := []string{"2026", "11", "10", "12"}, []rune{'-', '-', '~'}
	if m := signFacts(labels, signs, 0); len(m.Parts) != 3 || m.Dots != 2 || m.Parts[2] != "10" {
		t.Fatal(m)
	}
	if m := signFacts(labels, signs, 1); len(m.Parts) != 3 || m.Ext != "10" {
		t.Fatal(m)
	}
	if m := signFacts(labels, signs, 2); len(m.Parts) != 2 || m.Parts[0] != "10" || m.Parts[1] != "12" {
		t.Fatal(m)
	}
}
