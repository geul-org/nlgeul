//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReadOnto — a closing sign with a particle written onto it is read as two bundles: the sign (its closing, and its letter alone when it is unsure), then the particle with no stem (+를[목적격:를]) at verdict 1, beside the particle's letters as a new word; the particle's preference is the old one less the sign's own; the signs after the letters come last; a trace has the head's lines once
package token

import (
	"math"
	"testing"
)

func TestReadOnto(t *testing.T) {
	ix, _ := newIndex(Options{})
	u := unit{text: ")를.", head: 1, name: pairName, reads: []pairing{{shutName, 1}}}
	got, notes := readOnto(ix, u, false)
	if len(got) != 3 || got[0].Text != ")" || got[1].Text != "를" || got[2].Text != "." || len(notes) != 0 {
		t.Fatal(got, notes)
	}
	particle := got[1].Branches[0]
	if !particle.Words[0].Attach || particle.Words[0].Parts[0].Text != "를" || particle.Verdict != 1 || math.Abs(particle.Pref+1.42) > 0.01 {
		t.Fatal("the particle", particle)
	}
	whole := got[1].Branches[len(got[1].Branches)-1]
	if whole.Words[0].Attach || whole.Words[0].Parts[0].Text != "를" || len(whole.Words[0].Parts) != 1 {
		t.Fatal("the letters as a new word", whole)
	}
	if len(got[0].Branches) != 1 || got[0].Branches[0].Words[0].Parts[0].Name.Kind != shutName {
		t.Fatal("the sign", got[0])
	}
	only, _ := readOnto(ix, unit{text: ")?", head: 1, name: pairName, reads: []pairing{{shutName, 1}}}, false)
	if len(only) != 2 || only[0].Text != ")" || only[1].Text != "?" {
		t.Fatal("signs only after the sign", only)
	}
	traced, notes := readOnto(ix, unit{text: ")를", head: 1, name: pairName, reads: []pairing{{shutName, 1}}, notes: []string{"head"}}, true)
	count := 0
	for _, n := range notes {
		if n == "head" {
			count++
		}
	}
	if len(traced) != 2 || count != 1 {
		t.Fatal("the trace", count, notes)
	}
	middle, _ := readOnto(ix, unit{text: ")밥·국", head: 1, name: pairName, reads: []pairing{{shutName, 1}}}, false)
	found := false
	for _, br := range middle[1].Branches {
		found = found || len(br.Words) == 3 && br.Words[1].Parts[0].Name.Kind == "가운뎃점"
	}
	if !found {
		t.Fatal("a middle dot in the letters after the sign", middle)
	}
}
