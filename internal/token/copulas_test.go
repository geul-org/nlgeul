//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCopulas — 곰이다 reads 이 as the copula stem before 다; 학교다 a dropped stem written (이); 누군 keeps the stem's syllable; 이다 alone and a tail without a closing give nothing; after a unit the copula takes the unit with it; before a closing of droppableKnown (네) the dropped stem is read only with a word table and asks for a nominal of it, while the stem said (곰이네) and the old closings (다) ask for nothing; a unit read as attaching has a stem before 이; end to end 거네 is 거+(이)+네 with a table and not without, a word no table has takes no dropped stem before 네 but does before 다
package token

import (
	"strings"
	"testing"
)

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
	ne := &entry{name: "맺음:네"}
	ws := wordsScan("거네", layerCost)
	if got = ws.copulas([]tail{{start: 1, pieces: []piece{{text: "네"}}, close: ne}}); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != ws.ix.nominal || got[0].gate == nil {
		t.Fatal("a dropped stem before 네 asks for a nominal", got)
	}
	if got = scanOf("거네").copulas([]tail{{start: 1, pieces: []piece{{text: "네"}}, close: ne}}); len(got) != 0 {
		t.Fatal("before 네 no dropped stem without a word table", got)
	}
	if got = wordsScan("곰이네", layerCost).copulas([]tail{{start: 2, pieces: []piece{{text: "네"}}, close: ne}}); len(got) != 1 || got[0].pieces[0].text != "이" || got[0].gate != nil {
		t.Fatal("the stem said before 네 asks for nothing", got)
	}
	if got = wordsScan("학교다", layerCost).copulas([]tail{{start: 2, pieces: []piece{{text: "다"}}, close: da}}); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != nil {
		t.Fatal("before 다 a dropped stem asks for nothing", got)
	}
	lone := scanOf("이다")
	lone.attach = true
	if got = lone.copulas([]tail{{start: 1, pieces: []piece{{text: "다"}}, close: da}}); len(got) != 1 || got[0].start != 0 || got[0].pieces[0].names != copulaName {
		t.Fatal("a unit read as attaching has its stem in the unit before", got)
	}
	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	if got := strings.Join(r.Lines("거네", false), "\n"); !strings.Contains(got, "거[의존명사:거]+(이)[서술격:이(빠짐)]+네[맺음:네]") {
		t.Fatal("거네 with a word table\n" + got)
	}
	if got := strings.Join(r.Lines("고래네", false), "\n"); strings.Contains(got, "(이)[서술격:이(빠짐)]+네") {
		t.Fatal("고래, a word no table has, before 네\n" + got)
	}
	if got := strings.Join(r.Lines("고래다", false), "\n"); !strings.Contains(got, "고래+(이)[서술격:이(빠짐)]+다[맺음:다]") {
		t.Fatal("고래 before 다, as before\n" + got)
	}
	for _, in := range []string{"거네", "고래네"} {
		if got := strings.Join(Lines(in, false), "\n"); strings.Contains(got, "(이)[서술격:이(빠짐)]+네") {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
}
