//ff:func feature=token type=reader control=sequence
//ff:what TestLoneDigit — a single Sino-Korean digit is lone only with no word table; a native numeral, a greater number and a longer run are not
package token

import "testing"

func TestLoneDigit(t *testing.T) {
	ix := scanOf("x").ix
	if !ix.loneDigit("이") || ix.loneDigit("두") || ix.loneDigit("백") || ix.loneDigit("이십") {
		t.Fatal("no word table")
	}
	if wordsScan("x", layerCost).ix.loneDigit("이") {
		t.Fatal("with a word table")
	}
}
