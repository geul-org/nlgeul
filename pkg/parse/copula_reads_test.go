//ff:func feature=parse type=reader control=sequence
//ff:what TestCopulaReads — a tail read after the said stem (곰이다) and after the dropped one (학교다, written (이)); a stem the graph does not take gives nothing (학교서); a particle read after the tail is told to the graph (문제기도, read once more after the suffix 제 of 문제, which the gate of a suffix judges later); a fused piece keeps the stem's syllable (누군지); right after a sign that closes a pair a dropped stem is read before 다 ((가)다), after a Latin letter it is not (abc다); with no letter before the stem the unit must attach (이다); before a closing of droppableKnown the dropped stem asks for a nominal (거네), and a stem dropped right after a syllable 이 asks for one whatever follows (고양이다), and so does a stem dropped under a lone ㄴ that asked for a predicate's stem (문젠); after a naming closing the tail is read once more with the closing before the stem (먹음이다)
package parse

import "testing"

func TestCopulaReads(t *testing.T) {
	da := &entry{name: "맺음:다"}
	one := tail{start: 2, pieces: []piece{{text: "다", names: "맺음:다"}}, close: da, noYo: true}
	sc := scanOf("곰이다")
	got := sc.copulaReads(one, 1, sc.ix.copula, false)
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].text != "이" || got[0].pieces[0].unseen || !got[0].noYo || got[0].close != da || got[0].gate != nil {
		t.Fatal(got)
	}
	sc = scanOf("학교다")
	if got = sc.copulaReads(one, 2, sc.ix.dropped, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || !got[0].pieces[0].unseen || got[0].gate != nil {
		t.Fatal(got)
	}
	sc = scanOf("학교서")
	if got = sc.copulaReads(tail{start: 2, pieces: []piece{{text: "서", names: "맺음:서"}}, close: &entry{name: "맺음:서"}}, 2, sc.ix.dropped, false); got != nil {
		t.Fatal("no dropped stem before 서", got)
	}
	naming := tail{start: 2, pieces: []piece{{text: "기", names: "이름:기"}}, close: &entry{name: "이름:기"}}
	ch := wordsScan("문제기도", layerCost)
	if got = ch.copulaReads(naming, 2, ch.ix.dropped, false); got != nil {
		t.Fatal("no dropped stem before 기 with no particle after it", got)
	}
	if got = ch.copulaReads(naming, 2, ch.ix.dropped, true); len(got) != 2 || got[0].gate != ch.ix.nominal || got[1].pieces[0].text != "제" {
		t.Fatal("a dropped stem before 기 with a particle after it, and once more after the suffix 제 of the word 문제 (its gate judges it later)", got)
	}
	sc = scanOf("누군지")
	if got = sc.copulaReads(tail{start: 2, pieces: []piece{{text: "(ㄴ)지", names: "맺음:ㄴ지"}}, base: '구', close: &entry{name: "맺음:ㄴ지"}}, 2, sc.ix.dropped, false); len(got) != 1 || got[0].base != '구' {
		t.Fatal("the stem's syllable is kept", got)
	}
	lone := scanOf("이다")
	head := tail{start: 1, pieces: one.pieces, close: da}
	if got = lone.copulaReads(head, 0, lone.ix.copula, false); got != nil {
		t.Fatal("no stem before the copula", got)
	}
	lone.attach = true
	if got = lone.copulaReads(head, 0, lone.ix.copula, false); len(got) != 1 || got[0].start != 0 {
		t.Fatal("a unit read as attaching", got)
	}
	sign := wordsScan("(가)다", layerCost)
	if got = sign.copulaReads(tail{start: 3, pieces: one.pieces, close: da}, 3, sign.ix.dropped, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != nil {
		t.Fatal("a dropped stem before 다 right after a sign that closes a pair", got)
	}
	run := wordsScan("abc다", layerCost)
	if got = run.copulaReads(tail{start: 3, pieces: one.pieces, close: da}, 3, run.ix.dropped, false); len(got) != 1 || got[0].pieces[0].text != "(이)" || got[0].gate != nil {
		t.Fatal("a dropped stem before 다 after a Latin letter, asking for no nominal", got)
	}
	named := wordsScan("먹음이다", layerCost)
	if got = named.copulaReads(tail{start: 3, pieces: one.pieces, close: da}, 2, named.ix.copula, false); len(got) < 2 || got[1].start != 1 || got[1].pieces[0].names != "이름:음" || got[1].pieces[1].names != copulaName {
		t.Fatal("the copula read once more after the naming closing 음", got)
	}
	short := wordsScan("문젠", layerCost)
	if got = short.copulaReads(tail{start: 2, pieces: []piece{{text: "(ㄴ)", names: "꾸밈:ㄴ"}}, base: '제', close: &entry{name: "꾸밈:ㄴ"}, gate: short.ix.pred}, 2, short.ix.dropped, false); len(got) != 1 || got[0].base != '제' || got[0].gate != short.ix.nominal {
		t.Fatal("a stem dropped under a lone ㄴ that asked for a predicate's stem asks for a nominal", got)
	}
	ws := wordsScan("거네", layerCost)
	if got = ws.copulaReads(tail{start: 1, pieces: []piece{{text: "네", names: "맺음:네"}}, close: &entry{name: "맺음:네"}}, 1, ws.ix.dropped, false); len(got) != 1 || got[0].gate != ws.ix.nominal || got[0].gate == nil {
		t.Fatal("a dropped stem before 네 asks for a nominal", got)
	}
	cat := wordsScan("고양이다", layerCost)
	if got = cat.copulaReads(tail{start: 3, pieces: one.pieces, close: da}, 3, cat.ix.dropped, false); len(got) != 1 || got[0].start != 3 || got[0].pieces[0].text != "(이)" || got[0].gate != cat.ix.nominal {
		t.Fatal("a stem dropped right after a syllable 이 asks for a nominal", got)
	}
}
