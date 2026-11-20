//ff:func feature=parse type=reader control=sequence
//ff:what TestWholeWord — a lone Hangul numeral and a unit are a word together when the word table has them, or the closed set of fused forms; never without a word table, nor for a longer numeral the table has no word of (이십건) or for a Latin letter
package parse

import "testing"

func TestWholeWord(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if !ix.wholeWord("세", "대") || !ix.wholeWord("이", "건") || ix.wholeWord("삼", "대") {
		t.Fatal("word table")
	}
	if ix.wholeWord("이십", "건") || ix.wholeWord("k", "건") {
		t.Fatal("no word of the table, or no Hangul numeral")
	}
	if scanOf("x").ix.wholeWord("이", "건") {
		t.Fatal("no word table")
	}
}
