//ff:func feature=rule type=helper control=sequence
//ff:what TestAroundFactsAdnominal — the facts of the adnominal and the word after it: each candidate of the word right before is an adnominal (AroundAfterAdnominal) and no adnominal may stand right before the word right after (AroundBeforeUnmodifiable), each alone and both together; the farther words do not count, and the candidate's own facts are not the neighbours'
package rule

import "testing"

func TestAroundFactsAdnominal(t *testing.T) {
	if got := aroundFacts(Around{Before: []Near{{AdnominalAll: true}}}); got != AroundAfterAdnominal {
		t.Fatal("after an adnominal", got)
	}
	if got := aroundFacts(Around{After: []Near{{Unmodifiable: true}}}); got != AroundBeforeUnmodifiable {
		t.Fatal("before a word no adnominal modifies", got)
	}
	both := Around{Before: []Near{{AdnominalAll: true}}, After: []Near{{Unmodifiable: true}}}
	if got := aroundFacts(both); got != AroundAfterAdnominal|AroundBeforeUnmodifiable {
		t.Fatal("both", got)
	}
	far := Around{Before: []Near{{}, {AdnominalAll: true}}, After: []Near{{}, {Unmodifiable: true}}}
	if got := aroundFacts(far); got != 0 {
		t.Fatal("only the nearest words count", got)
	}
	self := Around{Self: Near{AdnominalAll: true, Unmodifiable: true}}
	if got := aroundFacts(self); got != 0 {
		t.Fatal("the candidate's own word is no neighbour", got)
	}
}
