//ff:func feature=parse type=reader control=sequence
//ff:what TestReadsKnown — letters that read as a known word with markers (씨 before 가) or with none (학교), or as a run (A after Hangul), read known; letters that read only as a new word (곰곰, 곰곰이), as a restored syllable that may not be put back (싸, 쓰 with 아), or whose only tail takes all of them (를 read as a particle leaves no stem) do not; 가, where the particle takes all of the letters, reads known as the verb's stem with 아 fused into it
package parse

import "testing"

func TestReadsKnown(t *testing.T) {
	title := wordsScan("A씨가", layerCost)
	title.ix.lex["씨"] = word{kinds: kindNominal, count: 500, by: [6]int{500}}
	if !title.readsKnown(1, 3) {
		t.Fatal("씨가: the noun 씨 before 가")
	}
	if !wordsScan("A학교", layerCost).readsKnown(1, 3) {
		t.Fatal("학교: a word of the table with no marker")
	}
	if !wordsScan("가A", layerCost).readsKnown(1, 2) {
		t.Fatal("A: a run")
	}
	if wordsScan("A곰곰", layerCost).readsKnown(1, 3) {
		t.Fatal("곰곰: a new word only")
	}
	if wordsScan("A곰곰이", layerCost).readsKnown(1, 4) {
		t.Fatal("곰곰이: a new word only, with markers after it or not")
	}
	if wordsScan("A싸", layerCost).readsKnown(1, 2) {
		t.Fatal("싸: 쓰 with 아 may not be put back, and 싸 is no word")
	}
	if !wordsScan("A가", layerCost).readsKnown(1, 2) {
		t.Fatal("가: the particle takes all of the letters, but the verb's stem 가 stands with 아 fused into it")
	}
	if wordsScan("A를", layerCost).readsKnown(1, 2) {
		t.Fatal("를: the particle takes all of the letters, and 를 is no word")
	}
}
