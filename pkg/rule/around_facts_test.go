//ff:func feature=rule type=helper control=sequence
//ff:what TestAroundFacts — the facts of a candidate in its sentence as flags in the order of their constants: none for an empty one, each alone, and all three together; the words before and after count by the nearest only
package rule

import "testing"

func TestAroundFacts(t *testing.T) {
	if got := aroundFacts(Around{}); got != 0 {
		t.Fatal("nothing", got)
	}
	if got := aroundFacts(Around{Self: Near{Glued: true}}); got != AroundGlued {
		t.Fatal("glued", got)
	}
	if got := aroundFacts(Around{Before: []Near{{Nominal: true}}}); got != AroundAfterNominal {
		t.Fatal("after a nominal", got)
	}
	if got := aroundFacts(Around{After: []Near{{Modifiable: true}}}); got != AroundBeforeModifiable {
		t.Fatal("before a modifiable word", got)
	}
	all := Around{Self: Near{Glued: true}, Before: []Near{{Nominal: true}}, After: []Near{{Modifiable: true}}}
	if got := aroundFacts(all); got != AroundGlued|AroundAfterNominal|AroundBeforeModifiable {
		t.Fatal("all", got)
	}
	far := Around{Before: []Near{{}, {Nominal: true}}, After: []Near{{}, {Modifiable: true}}}
	if got := aroundFacts(far); got != 0 {
		t.Fatal("only the nearest words count", got)
	}
}
