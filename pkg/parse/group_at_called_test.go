//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestGroupAtCalled — the suffix 이 after a name stands at the end of its unit's letters only when the scan is told that a comma or an exclamation mark follows them (called), and then with no piece after it; inside the letters and before another particle the fact is not set, and the subject particle stands either way
package parse

import "testing"

func TestGroupAtCalled(t *testing.T) {
	names := func(pieces map[string]piece) map[string]bool {
		out := map[string]bool{}
		for _, pc := range pieces {
			out[pc.names] = true
		}
		return out
	}
	sc := scanOf("영숙이")
	plain, _, _ := sc.groupAt(classParticle, 3, -1, false, false, false)
	sc.called = true
	called, from, starts := sc.groupAt(classParticle, 3, -1, false, false, false)
	if got := names(plain); got[nameSuffix] || !got["주격:이"] {
		t.Fatal(plain)
	}
	if got := names(called); !got[nameSuffix] || !got["주격:이"] || !got["보격:이"] {
		t.Fatal(called)
	}
	for key, e := range from {
		if e.name == nameSuffix && (!e.called || e.gate == nil || starts[key] != 2) {
			t.Fatal(e, starts)
		}
	}
	before, from, _ := sc.groupAt(classParticle, 3, sc.ix.byKey['가'][0].slot, false, false, false)
	if !names(before)[nameSuffix] {
		t.Fatal(before)
	}
	for _, e := range from {
		if e.name == nameSuffix && !e.inner {
			t.Fatal(e)
		}
	}
	sc = scanOf("영숙이가")
	sc.called = true
	inside, _, _ := sc.groupAt(classParticle, 3, -1, false, false, false)
	if names(inside)[nameSuffix] {
		t.Fatal(inside)
	}
}
