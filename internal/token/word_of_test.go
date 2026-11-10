//ff:func feature=token type=reader control=sequence
//ff:what TestWordOf — a new word with its marker: the low verdict, the stem and marker printed with the graph name; a closed-set word as stem has verdict 1; a stem a piece is fused into is printed and looked up as it was before fusing; with a word table a stem is known only when its kind takes what follows, and its preference comes from the table's count less the cost of structure; a stem read as a number has the preference the tail gives it
package token

import (
	"math"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestWordOf(t *testing.T) {
	sc := scanOf("곰이것")
	e := sc.wordOf(0, 2, tail{start: 1, pieces: []piece{{"이", "주격:이", 1, -1.2, 0, false}}})
	if e.marked != "곰+이[주격:이]" || e.names != "새 낱말 + 주격:이" || e.verdict != sc.ix.unknown || e.end != 2 {
		t.Fatal(e)
	}
	if e = sc.wordOf(2, 3, tail{start: 3}); e.marked != "것[의존명사:것]" || e.verdict != 1 {
		t.Fatal(e)
	}
	ws := wordsScan("고양이가마신", layerCost)
	e = ws.wordOf(0, 4, tail{start: 3, pieces: []piece{{"가", "주격:가", 1, -1.2, rule.TailCase, false}}})
	if e.marked != "고양이[열쇠:체언]+가[주격:가]" || e.verdict <= ws.ix.unknown || math.Abs(e.pref-(math.Log10(1000/ws.ix.lexTotal)-1.2-2*costPiece)) > 1e-9 {
		t.Fatal(e)
	}
	if e = ws.wordOf(4, 6, tail{start: 6, pieces: []piece{{"다", "맺음:다", 1, -4, rule.TailPred, false}}}); e.verdict != ws.ix.unknown {
		t.Fatal("a noun before a closing counted as known", e)
	}
	rare := wordsScan("of", layerCost)
	if e = rare.wordOf(0, 2, tail{start: 2}); e.verdict != rare.ix.unknown || e.marked != "of[열쇠:홀로]" {
		t.Fatal("a rarely seen known word below a new word", e)
	}
	at := scanOf("task.py에")
	if e = at.wordOf(0, 8, tail{start: 7, pieces: []piece{{"에", "부사격:에", 1, -1.2, rule.TailCase, false}}}); e.marked != "task.py[파일명|도메인]+에[부사격:에]" || e.verdict != 1 || math.Abs(e.pref-(math.Log10(1.0/syllables)-1.2)) > 1e-9 {
		t.Fatal(e)
	}
	sc = scanOf("마신다")
	e = sc.wordOf(0, 3, tail{start: 2, base: '시', pieces: []piece{{"(ㄴ)다", "맺음:ㄴ다", 1, -4, 0, false}}})
	if e.marked != "마시+(ㄴ)다[맺음:ㄴ다]" || e.end != 3 {
		t.Fatal(e)
	}
	if e = scanOf("봐").wordOf(0, 1, tail{start: 1, base: '보', pieces: []piece{{"(아)", "맺음:어(합침)", 1, 0, 0, false}}}); e.marked != "보[보조 용언:보]+(아)[맺음:어(합침)]" || e.verdict != 1 {
		t.Fatal(e)
	}
	num := scanOf("세명")
	if e := num.wordOf(0, 2, tail{start: 1, as: numName, weight: -1.5, pieces: []piece{{text: "명", names: "단위:명", verdict: 1}}}); e.marked != "세[수]+명[단위:명]" || e.pref != -1.5 || e.verdict != 1 {
		t.Fatal(e)
	}
}
