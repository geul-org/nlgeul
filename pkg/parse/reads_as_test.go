//ff:func feature=parse type=reader control=sequence
//ff:what TestReadsAs — the letters from a place to another read as a known word the test accepts, with no marker or with markers that end there: 손으로 as the noun 손 with 으로, 와 as the verb 오 with 아 fused in; not when the test wants another kind or another stem, not as a new word (곰곰), and not with a last syllable that may not be restored so
package parse

import "testing"

func TestReadsAs(t *testing.T) {
	sc := wordsScan("손으로", layerCost)
	sc.ix.lex["손"] = word{kinds: kindNominal, count: 500, by: [6]int{500}}
	nominal := func(r stemCandidate) bool { return r.kind&kindNominal != 0 }
	verbal := func(r stemCandidate) bool { return r.kind&(kindVerb|kindAdj) != 0 }
	if !sc.readsAs(0, 3, nominal) {
		t.Fatal("손으로 as a noun")
	}
	if sc.readsAs(0, 3, verbal) {
		t.Fatal("손으로 as a predicate")
	}
	if sc.readsAs(0, 3, func(r stemCandidate) bool { return r.stem == "거" }) {
		t.Fatal("손으로 as 거")
	}
	come := wordsScan("와", layerCost)
	come.ix.lex["오"] = word{kinds: kindVerb, count: 10000, by: [6]int{0, 10000}}
	if !come.readsAs(0, 1, verbal) {
		t.Fatal("와 as 오+아")
	}
	if wordsScan("곰곰", layerCost).readsAs(0, 2, func(stemCandidate) bool { return true }) {
		t.Fatal("a new word does not count")
	}
	if wordsScan("싸", layerCost).readsAs(0, 1, func(r stemCandidate) bool { return r.stem == "쓰" }) {
		t.Fatal("싸 of one syllable: 아 put back under ㅡ does not stand (the syllable is not restored so)")
	}
}
