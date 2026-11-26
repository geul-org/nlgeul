//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCopulas — 곰이다 reads 이 as the copula stem before 다; 학교다 a dropped stem written (이); 누군 keeps the stem's syllable; 이다 alone and a tail without a closing give nothing; after a unit the copula takes the unit with it; before a closing of droppableKnown (네) the dropped stem is read only with a word table and asks for a nominal of it, before a quoting closing of droppableNew (라는) for a nominal of it or a new word, while the stem said (곰이네) and the old closings (다) ask for nothing; after a letter that is no Hangul syllable a dropped stem stands before 야 asking for no nominal, and not before 다; before 기 it stands only while a particle is read after it (chained), asking for a nominal; the dropped stem's graph turns down a closing it is not dropped before (서, with a word table too), a slot before the closing, a consonant before it, and without a word table 야 and 기; a unit read as attaching has a stem before 이; end to end 거네 is 거+(이)+네 with a table and not without, a word no table has takes no dropped stem before 네 but does before the quoting 라는 and before 다; the copula's stem, said or dropped, is not read before 하 of 어하다 written onto its 어/아 form (학교이어하는, 문제여하는), while 하 of 어하다 after a verb or an adjective (예뻐하는, 몰라하는) and the copula before another closing (학교이어서, 학교여서) stay, with a word table and without
package parse

import (
	"strings"
	"testing"
)

func TestCopulas(t *testing.T) {
	da := &entry{name: "맺음:다"}
	sc := scanOf("곰이다")
	got := sc.copulas([]tail{{start: 2, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da, noYo: true}, {start: 2, pieces: []piece{{text: "다", names: "맺음:다"}}}}, false)
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].names != copulaName || got[0].pieces[0].text != "이" || got[0].pieces[1].text != "다" || !got[0].noYo || got[0].close != da {
		t.Fatal(got)
	}
	got = scanOf("학교다").copulas([]tail{{start: 2, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false)
	if len(got) != 1 || got[0].start != 2 || got[0].pieces[0].text != "(이)" || !got[0].pieces[0].unseen || got[0].base != 0 {
		t.Fatal(got)
	}
	got = scanOf("누군").copulas([]tail{{start: 2, pieces: []piece{{text: "(ㄴ)", names: "꾸밈:ㄴ"}}, base: '구', close: &entry{name: "꾸밈:ㄴ"}}}, false)
	if len(got) != 1 || got[0].base != '구' || got[0].start != 2 {
		t.Fatal(got)
	}
	if got = scanOf("이다").copulas([]tail{{start: 1, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false); len(got) != 0 {
		t.Fatal("no stem before the copula", got)
	}
	got = scanOf("3개다").copulas([]tail{{start: 2, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false)
	if len(got) != 2 || got[1].start != 1 || got[1].pieces[0].names != "단위:개" {
		t.Fatal(got)
	}
	ne := &entry{name: "맺음:네"}
	ws := wordsScan("거네", layerCost)
	if got = ws.copulas([]tail{{start: 1, pieces: []piece{{text: "네", names: "맺음:네"}}, close: ne}}, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != ws.ix.nominal || got[0].gate == nil {
		t.Fatal("a dropped stem before 네 asks for a nominal", got)
	}
	rane := &entry{name: "꾸밈:라는"}
	quoted := wordsScan("뚜루라는", layerCost)
	if got = quoted.copulas([]tail{{start: 2, pieces: []piece{{text: "라는", names: "꾸밈:라는"}}, close: rane}}, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != quoted.ix.nominalNew || got[0].gate == nil {
		t.Fatal("a dropped stem before the quoting 라는 asks for a nominal or a new word", got)
	}
	cat := wordsScan("고양이다", layerCost)
	if got = cat.copulas([]tail{{start: 3, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false); len(got) != 2 || got[0].start != 2 || got[0].pieces[0].text != "이" || got[1].start != 3 || got[1].pieces[0].text != "(이)" || got[1].gate != cat.ix.nominal {
		t.Fatal("after a syllable 이 the stem said and the stem dropped", got)
	}
	honor := []tail{{start: 2, pieces: []piece{{text: "시", names: "선어말:시"}, {text: "다", names: "맺음:다"}}, close: da}}
	polite := wordsScan("어디시다", layerCost)
	if got = polite.copulas(honor, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].pieces[1].names != "선어말:시" || got[0].gate != polite.ix.nominal || got[0].gate == nil {
		t.Fatal("a dropped stem before the slot 시 asks for a nominal", got)
	}
	if got = scanOf("어디시다").copulas(honor, false); len(got) != 0 {
		t.Fatal("before the slot 시 no dropped stem without a word table", got)
	}
	if got = scanOf("거네").copulas([]tail{{start: 1, pieces: []piece{{text: "네", names: "맺음:네"}}, close: ne}}, false); len(got) != 0 {
		t.Fatal("before 네 no dropped stem without a word table", got)
	}
	if got = wordsScan("곰이네", layerCost).copulas([]tail{{start: 2, pieces: []piece{{text: "네", names: "맺음:네"}}, close: ne}}, false); len(got) != 2 || got[0].pieces[0].text != "이" || got[0].gate != nil || got[1].pieces[0].text != "(이)" {
		t.Fatal("the stem said before 네 asks for nothing", got)
	}
	if got = wordsScan("학교다", layerCost).copulas([]tail{{start: 2, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != nil {
		t.Fatal("before 다 a dropped stem asks for nothing", got)
	}
	ya := &entry{name: "맺음:야"}
	run := wordsScan("abc야", layerCost)
	if got = run.copulas([]tail{{start: 3, pieces: []piece{{text: "야", names: "맺음:야"}}, close: ya}}, false); len(got) != 1 || got[0].start != 3 || got[0].pieces[0].text != "(이)" || got[0].pieces[0].names != droppedName || got[0].gate != nil {
		t.Fatal("after a letter that is no Hangul syllable a dropped stem before 야 asks for no nominal", got)
	}
	if got = wordsScan("abc다", layerCost).copulas([]tail{{start: 3, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false); len(got) != 0 {
		t.Fatal("after a Latin letter no dropped stem before 다", got)
	}
	shut := wordsScan("(가)다", layerCost)
	if got = shut.copulas([]tail{{start: 3, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false); len(got) != 1 || got[0].start != 3 || got[0].pieces[0].text != "(이)" || got[0].gate != nil {
		t.Fatal("right after a sign that closes a pair a dropped stem before 다, asking for no nominal", got)
	}
	naming := []tail{{start: 2, pieces: []piece{{text: "기", names: "이름:기"}}, close: &entry{name: "이름:기"}}}
	ch := wordsScan("문제기도", layerCost)
	if got = ch.copulas(naming, false); len(got) != 0 {
		t.Fatal("no dropped stem before 기 when no particle is read after it", got)
	}
	if got = ch.copulas(naming, true); len(got) != 1 || got[0].start != 2 || got[0].pieces[0].text != "(이)" || got[0].pieces[1].names != "이름:기" || got[0].gate != ch.ix.nominal {
		t.Fatal("a dropped stem before 기 with a particle read after it, asking for a nominal", got)
	}
	seo, slot := []piece{{text: "서", names: "맺음:서"}}, []piece{{text: "었", names: "선어말:었"}, {text: "다", names: "맺음:다"}}
	for _, c := range []struct {
		sc      *scan
		t       tail
		chained bool
	}{
		{scanOf("학교서"), tail{start: 2, pieces: seo, close: &entry{name: "맺음:서"}}, false},
		{wordsScan("학교서", layerCost), tail{start: 2, pieces: seo, close: &entry{name: "맺음:서"}}, false},
		{scanOf("학교었다"), tail{start: 2, pieces: slot, close: da}, false},
		{scanOf("곰고"), tail{start: 1, pieces: []piece{{text: "고", names: "맺음:고"}}, close: &entry{name: "맺음:고"}}, false},
		{scanOf("abc야"), tail{start: 3, pieces: []piece{{text: "야", names: "맺음:야"}}, close: ya}, false},
		{scanOf("abc다"), tail{start: 3, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}, false},
		{wordsScan("(가다", layerCost), tail{start: 1, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}, false},
		{scanOf("문제기도"), naming[0], true},
	} {
		if got = c.sc.copulas([]tail{c.t}, c.chained); len(got) != 0 {
			t.Errorf("%s: the dropped stem's graph takes no stem before %s here", string(c.sc.s), c.t.pieces[0].names)
		}
	}
	lone := scanOf("이다")
	lone.attach = true
	if got = lone.copulas([]tail{{start: 1, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da}}, false); len(got) != 1 || got[0].start != 0 || got[0].pieces[0].names != copulaName {
		t.Fatal("a unit read as attaching has its stem in the unit before", got)
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	if got := streamText(r, "거네"); !strings.Contains(got, "거[의존명사:거]+(이)[서술격:이(빠짐)]+네[맺음:네]") {
		t.Fatal("거네 with a word table\n" + got)
	}
	if got := streamText(r, "고래네"); strings.Contains(got, "(이)[서술격:이(빠짐)]+네") {
		t.Fatal("고래, a word no table has, before 네\n" + got)
	}
	if got := streamText(r, "뚜루라는"); !strings.Contains(got, "뚜루+(이)[서술격:이(빠짐)]+라는[꾸밈:라는]") {
		t.Fatal("뚜루, a word no table has, before the quoting 라는\n" + got)
	}
	if got := streamText(r, "고래다"); !strings.Contains(got, "고래+(이)[서술격:이(빠짐)]+다[맺음:다]") {
		t.Fatal("고래 before 다, as before\n" + got)
	}
	bare, _ := New(Options{})
	for _, in := range []string{"거네", "고래네"} {
		if got := streamText(bare, in); strings.Contains(got, "(이)[서술격:이(빠짐)]+네") {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
	plain, _ := New(Options{})
	for _, c := range []struct {
		in, want string
		dead     bool
	}{
		{"학교이어하는", "보조:하", true}, {"문제여하는", "보조:하", true},
		{"학교이어서", "이[서술격:이]+어서[맺음:어서]", false}, {"학교여서", "이[서술격:이]+(어)서[맺음:어서(합침)]", false},
	} {
		got, withTable := streamText(plain, c.in), streamText(r, c.in)
		if strings.Contains(got, c.want) == c.dead || strings.Contains(withTable, c.want) == c.dead {
			t.Errorf("%s: no auxiliary 하 right after the copula's 어/아, the copula before 어서 stays (dead: %v)\n%s\n%s", c.in, c.dead, got, withTable)
		}
	}
	for in, want := range map[string]string{"예뻐하는": "예쁘[열쇠:형용사]+(어)[맺음:어(합침)]+하[보조:하]+는[꾸밈:는]", "몰라하는": "모르[열쇠:동사]+(아)[맺음:어(합침)]+하[보조:하]+는[꾸밈:는]"} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
}
