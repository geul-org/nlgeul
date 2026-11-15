//ff:func feature=token type=reader control=sequence
//ff:what TestApartWord — letters that read only as a title of article 48 (씨 before 가, 님 as a whole) or a dependent noun of article 42 (측 before 이) are such a word; not so when they also read as a verb (내 before 는), when they are cut in two ways (양 before 의, and 양의), when the word is none of the sets (팝), when a tail takes all of the letters (가 read as a particle leaves no stem, and the stem 가 is a verb's), nor when the only reading is a restored syllable that may not be put back (싸, 쓰 with 아)
package token

import "testing"

func TestApartWord(t *testing.T) {
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}

	title := wordsScan("A씨가", layerCost)
	title.ix.lex["씨"] = noun
	if !title.apartWord(1, 3) {
		t.Fatal("씨가: 씨, a title of article 48, before 가")
	}
	honor := wordsScan("PD님", layerCost)
	honor.ix.lex["님"] = noun
	if !honor.apartWord(2, 3) {
		t.Fatal("님 as a whole")
	}
	side := wordsScan("A측이", layerCost)
	side.ix.lex["측"] = noun
	if !side.apartWord(1, 3) {
		t.Fatal("측이: 측, a dependent noun of article 42, before 이")
	}

	verb := wordsScan("OO내는", layerCost)
	verb.ix.lex["내"] = word{kinds: kindNominal | kindVerb, count: 500, by: [6]int{250, 250}}
	if verb.apartWord(2, 4) {
		t.Fatal("내는: 내 reads as a verb too")
	}
	cut := wordsScan("C양의", layerCost)
	cut.ix.lex["양"] = noun
	cut.ix.lex["양의"] = noun
	if cut.apartWord(1, 3) {
		t.Fatal("양의: cut as 양+의 and read as 양의")
	}
	if !cut.apartWord(1, 2) {
		t.Fatal("양 alone is cut in one way only")
	}
	pop := wordsScan("K팝을", layerCost)
	pop.ix.lex["팝"] = noun
	if pop.apartWord(1, 3) {
		t.Fatal("팝을: 팝 is in none of the sets")
	}
	if wordsScan("A가", layerCost).apartWord(1, 2) {
		t.Fatal("가: the particle takes all of the letters, and the stem 가 is a verb's")
	}
	if wordsScan("A싸", layerCost).apartWord(1, 2) {
		t.Fatal("싸: 쓰 with 아 may not be put back, and 싸 is no word")
	}
}
