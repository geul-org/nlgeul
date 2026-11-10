//ff:func feature=token type=reader control=sequence
//ff:what TestWholeWord — a lone Hangul numeral and a unit are a word together when the word table has them, or the closed set of fused forms; never without a word table, for a longer numeral or for a Latin letter
package token

import "testing"

func TestWholeWord(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if !ix.wholeWord("세", "대") || !ix.wholeWord("이", "건") || ix.wholeWord("삼", "대") {
		t.Fatal("word table")
	}
	if ix.wholeWord("이십", "건") || ix.wholeWord("k", "건") {
		t.Fatal("not a lone Hangul numeral")
	}
	if scanOf("x").ix.wholeWord("이", "건") {
		t.Fatal("no word table")
	}
}
