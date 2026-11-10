//ff:func feature=token type=reader control=sequence
//ff:what TestRestoresAt — with a word table ㅕ comes from ㅣ+어 and from the stem's own ㅕ, under ㅆ too; without one, and for another vowel, only as restores gives
package token

import "testing"

func TestRestoresAt(t *testing.T) {
	ws := wordsScan("켜켰가", layerCost)
	if got := ws.restoresAt(1, leadVowel); len(got) != 2 || got[0].stem != '키' || got[1].stem != '켜' || got[1].piece != "어" {
		t.Fatal(got)
	}
	if got := ws.restoresAt(2, jongSS); len(got) != 2 || got[1].stem != '켜' || got[1].piece != "었" {
		t.Fatal(got)
	}
	if got := ws.restoresAt(3, leadVowel); len(got) != 2 || got[0].stem != '가' {
		t.Fatal("another vowel", got)
	}
	if got := scanOf("켜").restoresAt(1, leadVowel); len(got) != 1 {
		t.Fatal("no word table", got)
	}
}
