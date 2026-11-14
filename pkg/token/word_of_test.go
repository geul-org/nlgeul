//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWordOf — a new word with its marker: the low verdict, the stem and marker printed with the graph name; a closed-set word as stem has verdict 1; a stem a piece is fused into is printed and looked up as it was before fusing; with a word table a stem is known only when its kind takes what follows, a stem the table knows only as a nominal does not stand as a new word before a predicate tail either (verdict 0 or less), and its preference comes from the table's count less the cost of structure; a stem read as a number has the preference the tail gives it; the boundary given back to an auxiliary verb's stem after a boundary only when a predicate tail follows it and it is known there (not before particles or the copula, not alone, not as the unit's first word, not for another verb; with a word table 있 before an adjective's closing is no known word, and as a new word that ends in ㅆ before markers it does not stand)
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWordOf(t *testing.T) {
	sc := scanOf("곰이것")
	e := sc.wordOf(0, 2, tail{start: 1, pieces: []piece{{"이", "주격:이", 1, -1.2, 0, false, ""}}})
	if e.sig != "곰+이[주격:이]" || (Branch{Words: []Eojeol{e.word}}).Label() != "새 낱말 + 주격:이" || e.verdict != sc.ix.unknown || e.end != 2 {
		t.Fatal(e)
	}
	if e = sc.wordOf(2, 3, tail{start: 3}); e.sig != "것[의존명사:것]" || e.verdict != 1 {
		t.Fatal(e)
	}
	ws := wordsScan("고양이가마신", layerCost)
	e = ws.wordOf(0, 4, tail{start: 3, pieces: []piece{{"가", "주격:가", 1, -1.2, rule.TailCase, false, ""}}})
	if e.sig != "고양이[열쇠:체언]+가[주격:가]" || e.verdict <= ws.ix.unknown || math.Abs(e.pref-(math.Log10(1000/ws.ix.lexTotal)-1.2-costPiece)) > 1e-9 {
		t.Fatal(e)
	}
	if e = ws.wordOf(4, 6, tail{start: 6, pieces: []piece{{"다", "맺음:다", 1, -4, rule.TailPred, false, ""}}}); e.verdict > 0 || e.word.Parts[0].Name != (Name{}) {
		t.Fatal("a noun of the table before a closing stands neither as known nor as a new word", e)
	}
	rare := wordsScan("of", layerCost)
	if e = rare.wordOf(0, 2, tail{start: 2}); e.verdict <= rare.ix.unknown || e.verdict >= 1 || e.sig != "of[열쇠:홀로]" {
		t.Fatal("a rarely seen known word stands by its graph, above a new word", e)
	}
	at := scanOf("task.py에")
	two := at.stepsOf(0, 8, tail{start: 7, pieces: []piece{{"에", "부사격:에", 1, -1.2, rule.TailCase, false, ""}}})
	if len(two) != 2 || two[0].sig != "task.py[파일명]+에[부사격:에]" || two[1].sig != "task.py[도메인]+에[부사격:에]" {
		t.Fatal("an atom of two names is two candidates", two)
	}
	for _, e = range two {
		if e.verdict != 1 || math.Abs(e.pref-(math.Log10(1.0/syllables)-1.2-costPiece)) > 1e-9 {
			t.Fatal("of the same verdict and preference", e)
		}
	}
	sc = scanOf("마신다")
	e = sc.wordOf(0, 3, tail{start: 2, base: '시', pieces: []piece{{"(ㄴ)다", "맺음:ㄴ다", 1, -4, 0, false, ""}}})
	if e.sig != "마시+(ㄴ)다[맺음:ㄴ다]" || e.end != 3 {
		t.Fatal(e)
	}
	if e = scanOf("봐").wordOf(0, 1, tail{start: 1, base: '보', pieces: []piece{{"(아)", "맺음:어(합침)", 1, 0, rule.TailPred, false, ""}}}); e.sig != "보[보조 용언:보]+(아)[맺음:어(합침)]" || e.verdict != 1 {
		t.Fatal(e)
	}
	main := wordsScan("먹어보다", layerCost)
	if e = main.wordOf(0, 2, tail{start: 1, pieces: []piece{{"어", "맺음:어", 1, -1, rule.TailPred, false, ""}}}); !e.conn || e.back != 0 {
		t.Fatal("a word that ends in 어", e)
	}
	if e = main.wordOf(2, 4, tail{start: 3, pieces: []piece{{"다", "맺음:다", 1, -1, rule.TailPred, false, ""}}}); e.conn || e.back != costBoundary {
		t.Fatal("an auxiliary verb's stem after a boundary", e)
	}
	aux := wordsScan("열어두는", layerCost)
	if e = aux.wordOf(2, 4, tail{start: 3, pieces: []piece{{"는", "꾸밈:는", 1, -1, rule.TailVerb, false, ""}}}); e.back != costBoundary || e.sig != "두[보조 용언:두]+는[꾸밈:는]" {
		t.Fatal("an auxiliary verb's stem before a predicate tail", e)
	}
	if e = aux.wordOf(2, 4, tail{start: 3, pieces: []piece{{"는", "보조사:는", 1, -1, rule.TailAux, false, ""}}}); e.back != 0 || e.verdict <= 0 {
		t.Fatal("the same letters standing as a noun before a particle", e)
	}
	if e = aux.wordOf(2, 4, tail{start: 3, pieces: []piece{{"를", "목적격:를", 1, -1, rule.TailCase, false, ""}}}); e.back != 0 || e.verdict <= 0 {
		t.Fatal("the same letters standing as a noun before a case particle", e)
	}
	if e = aux.wordOf(2, 4, tail{start: 3, pieces: []piece{{"이다", "서술격:이다", 1, -1, rule.TailCopula, false, ""}}}); e.back != 0 || e.verdict <= 0 {
		t.Fatal("the same letters standing as a noun before the copula", e)
	}
	if e = aux.wordOf(2, 3, tail{start: 3}); e.back != 0 || e.verdict <= 0 {
		t.Fatal("an auxiliary verb's letters standing alone", e)
	}
	if e = aux.wordOf(2, 4, tail{start: 3, pieces: []piece{{"다", "맺음:다", 1, -1, rule.TailPred, false, ""}}}); e.back != costBoundary {
		t.Fatal("an auxiliary verb's stem before a closing", e)
	}
	first := wordsScan("두는", layerCost)
	if e = first.wordOf(0, 2, tail{start: 1, pieces: []piece{{"는", "꾸밈:는", 1, -1, rule.TailVerb, false, ""}}}); e.back != 0 || e.verdict <= 0 {
		t.Fatal("an auxiliary verb's stem as the unit's first word", e)
	}
	other := wordsScan("열어먹다", layerCost)
	if e = other.wordOf(2, 4, tail{start: 3, pieces: []piece{{"다", "맺음:다", 1, -1, rule.TailPred, false, ""}}}); e.back != 0 || e.verdict <= 0 {
		t.Fatal("a known verb that is no auxiliary, after a boundary", e)
	}
	adj := []piece{{"다", "맺음:다", 1, -1, rule.TailAdj, false, ""}}
	if e = wordsScan("열어있다", layerCost).wordOf(2, 4, tail{start: 3, pieces: adj}); e.back != 0 || e.verdict > 0 || e.word.Parts[0].Name != (Name{}) {
		t.Fatal("an auxiliary verb's letters that stand as no known word before an adjective's closing, nor as a new word (a stem in ㅆ)", e)
	}
	if e = scanOf("열어있다").wordOf(2, 4, tail{start: 3, pieces: adj}); e.back != costBoundary || e.verdict != 1 {
		t.Fatal("without a word table, an auxiliary verb's stem before an adjective's closing", e)
	}
	if e = scanOf("열어두는").wordOf(2, 4, tail{start: 3, pieces: []piece{{"는", "보조사:는", 1, -1, rule.TailAux, false, ""}}}); e.back != 0 {
		t.Fatal("without a word table, an auxiliary verb's letters before a particle", e)
	}
	num := scanOf("세명")
	if e := num.wordOf(0, 2, tail{start: 1, as: numName, weight: -1.5, pieces: []piece{{text: "명", names: "단위:명", verdict: 1}}}); e.sig != "세[수]+명[단위:명]" || e.pref != -1.5-costPiece || e.verdict != 1 {
		t.Fatal(e)
	}
}
