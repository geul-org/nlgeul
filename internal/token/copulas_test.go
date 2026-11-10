//ff:func feature=token type=reader control=sequence
//ff:what TestCopulas — 곰이다 reads 이 as the copula stem before 다; 학교다 a dropped stem written (이); 누군 keeps the stem's syllable; 이다 alone and a tail without a closing give nothing; after a unit the copula takes the unit with it
package token

import "testing"

func TestCopulas(t *testing.T) {
	da := &entry{name: "맺음:다"}
	sc := scanOf("곰이다")
	got := sc.copulas([]tail{{start: 2, pieces: []piece{{text: "다"}}, close: da, noYo: true}, {start: 2, pieces: []piece{{text: "다"}}}})
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].names != copulaName || got[0].pieces[0].text != "이" || got[0].pieces[1].text != "다" || !got[0].noYo || got[0].close != da {
		t.Fatal(got)
	}
	got = scanOf("학교다").copulas([]tail{{start: 2, pieces: []piece{{text: "다"}}, close: da}})
	if len(got) != 1 || got[0].start != 2 || got[0].pieces[0].text != "(이)" || !got[0].pieces[0].unseen || got[0].base != 0 {
		t.Fatal(got)
	}
	got = scanOf("누군").copulas([]tail{{start: 2, pieces: []piece{{text: "(ㄴ)"}}, base: '구', close: &entry{name: "꾸밈:ㄴ"}}})
	if len(got) != 1 || got[0].base != '구' || got[0].start != 2 {
		t.Fatal(got)
	}
	if got = scanOf("이다").copulas([]tail{{start: 1, pieces: []piece{{text: "다"}}, close: da}}); len(got) != 0 {
		t.Fatal("no stem before the copula", got)
	}
	got = scanOf("3개다").copulas([]tail{{start: 2, pieces: []piece{{text: "다"}}, close: da}})
	if len(got) != 2 || got[1].start != 1 || got[1].pieces[0].names != "단위:개" {
		t.Fatal(got)
	}
}
