//ff:func feature=token type=reader control=sequence
//ff:what TestCopulasNumber — the copula's stem fused with the modifying ㄴ (인) right after a number (1인) is read with its verdict lowered by the rebuttal of numberAdnRebuttal, told by numEnds that a number ends right before it, and right after a Hangul syllable (곰인) it stands fully
package token

import "testing"

func TestCopulasNumber(t *testing.T) {
	adn := &entry{name: "꾸밈:ㄴ"}
	read := func(text string) []tail {
		return scanOf(text).copulas([]tail{{start: 2, pieces: []piece{{text: "(ㄴ)", names: "꾸밈:ㄴ"}}, base: '이', close: adn}}, false)
	}
	got := read("1인")
	if len(got) == 0 || got[0].start != 1 || got[0].pieces[0].names != copulaName || got[0].pieces[0].verdict <= 0.3 || got[0].pieces[0].verdict >= 0.34 {
		t.Fatal("right after a number", got)
	}
	got = read("곰인")
	if len(got) == 0 || got[0].pieces[0].names != copulaName || got[0].pieces[0].verdict != 1 {
		t.Fatal("right after a syllable", got)
	}
}
