//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestMarkerStem — a new word after a boundary whose letters read as markers there, or that starts with the link vowel 으 and from 으 on, or from after it, reads as markers up to the tail's end, is no stem (a stem a piece is fused into too, for the 으 rule only); the unit's first word, a stem whose letters are not so read, markers that end before the tail's end, a stem of another letter, a word the table knows, a stem a piece is fused into outside the 으 rule and a tail that names its stem are let be; so is every stem when there is no word table; the 으 left when ㄹ is taken off 을 (a stem of one syllable with 으 as its base) is the link vowel, while a fused stem of more than one syllable is judged by its first letter as written; end to end 먹었을걸 has no new word 으 before (ㄹ)걸
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerStem(t *testing.T) {
	if none := scanOf("곰에서"); none.markerStem(1, 3, tail{start: 2, pieces: []piece{{text: "서"}}}) || none.marks[2] != nil {
		t.Fatal("without a word table a new word's stem is not asked")
	}
	sc := wordsScan("곰에서", layerCost)
	sc.marks[2] = map[int]bool{1: true}
	seo := tail{start: 2, pieces: []piece{{text: "서", names: "부사격:서", verdict: 1, tail: rule.TailCase}}}
	if !sc.markerStem(1, 3, seo) {
		t.Fatal("에 after a boundary, read as a marker, before 서")
	}
	if sc.markerStem(0, 3, tail{start: 2, pieces: seo.pieces}) {
		t.Fatal("the unit's first word")
	}
	sc.marks[2] = map[int]bool{}
	if sc.markerStem(1, 3, seo) {
		t.Fatal("letters that do not read as markers")
	}
	sc.marks[2] = map[int]bool{1: true}
	if sc.markerStem(1, 3, tail{start: 2, base: '에', pieces: seo.pieces}) || sc.markerStem(1, 3, tail{start: 2, shown: "에", pieces: seo.pieces}) || sc.markerStem(1, 3, tail{start: 2, as: numName, pieces: seo.pieces}) || sc.markerStem(1, 3, tail{start: 2, only: "에", pieces: seo.pieces}) {
		t.Fatal("a tail with its own stem")
	}
	ws := wordsScan("곰서로", layerCost)
	ws.marks[2] = map[int]bool{1: true}
	if ws.markerStem(1, 3, tail{start: 2, pieces: []piece{{text: "로", names: "부사격:로", verdict: 1, tail: rule.TailCase}}}) {
		t.Fatal("a word the table knows (서 the noun)")
	}

	myeo := tail{start: 2, pieces: []piece{{text: "며", names: "맺음:며", verdict: 1, tail: rule.TailPred}}}
	link := wordsScan("곰으며", layerCost)
	if link.markerStem(1, 3, myeo) {
		t.Fatal("으 before 며 when nothing reads as markers")
	}
	link.marks[3] = map[int]bool{1: true}
	if !link.markerStem(1, 3, myeo) {
		t.Fatal("으 before 며: from the link vowel on the letters read as markers (으며)")
	}
	link.marks[3] = map[int]bool{2: true}
	if !link.markerStem(1, 3, myeo) {
		t.Fatal("으 before 며: after the link vowel the letters read as markers (며)")
	}
	if link.markerStem(1, 3, tail{start: 2, shown: "으", pieces: myeo.pieces}) || link.markerStem(1, 3, tail{start: 2, as: numName, pieces: myeo.pieces}) || link.markerStem(1, 3, tail{start: 2, only: "으", pieces: myeo.pieces}) {
		t.Fatal("a tail with its own stem, under the link vowel rule too")
	}
	link.marks[3] = map[int]bool{}
	link.marks[2] = map[int]bool{1: true}
	if !link.markerStem(1, 3, myeo) {
		t.Fatal("으 read as a marker by itself before 며, by the old rule")
	}
	if link.markerStem(1, 3, tail{start: 2, base: '으', pieces: myeo.pieces}) {
		t.Fatal("a stem a piece is fused into, when nothing up to the tail's end reads as markers")
	}
	head := wordsScan("으며", layerCost)
	head.marks[2] = map[int]bool{0: true, 1: true}
	if head.markerStem(0, 2, tail{start: 1, pieces: myeo.pieces}) {
		t.Fatal("the link vowel as the unit's first word")
	}
	if none := scanOf("곰으며"); none.markerStem(1, 3, myeo) {
		t.Fatal("without a word table the link vowel rule is not asked")
	}

	eu := wordsScan("곰으셨다", layerCost)
	da := tail{start: 3, pieces: []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}}
	if eu.markerStem(1, 4, da) {
		t.Fatal("으셨 before 다 when nothing reads 셨다 as markers")
	}
	eu.marks[3] = map[int]bool{2: true}
	if eu.markerStem(1, 4, da) {
		t.Fatal("markers that end before the tail's end do not count for the link vowel")
	}
	eu.marks[3] = nil
	eu.marks[4] = map[int]bool{2: true}
	if !eu.markerStem(1, 4, da) {
		t.Fatal("으셨 before 다: the link vowel and then a run of markers")
	}
	if eu.markerStem(2, 4, da) {
		t.Fatal("셨 before 다 does not start with the link vowel")
	}
	eu.marks[4] = map[int]bool{1: true}
	if !eu.markerStem(1, 4, da) {
		t.Fatal("a run of markers from the link vowel on")
	}
	if !eu.markerStem(1, 4, tail{start: 3, base: '시', pieces: da.pieces}) {
		t.Fatal("으시 with a piece fused into it, before 다")
	}

	known := wordsScan("곰으뜸이다", layerCost)
	cop := tail{start: 3, pieces: []piece{{text: "이다", names: "서술격:이다", verdict: 1, tail: rule.TailCopula}}}
	known.marks[5] = map[int]bool{2: true}
	if !known.markerStem(1, 5, cop) {
		t.Fatal("으뜸 before 이다 as a new word, when 뜸이다 reads as markers")
	}
	known.ix.lex["으뜸"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	if known.markerStem(1, 5, cop) {
		t.Fatal("a word the table knows that starts with the link vowel")
	}

	geol := []piece{{text: "(ㄹ)걸", names: "맺음:ㄹ걸", verdict: 1, tail: rule.TailPred}}
	eul := wordsScan("곰을걸", layerCost)
	if eul.markerStem(1, 3, tail{start: 2, base: '으', pieces: geol}) {
		t.Fatal("으 out of 을 before (ㄹ)걸 when nothing reads as markers")
	}
	eul.marks[3] = map[int]bool{1: true}
	if !eul.markerStem(1, 3, tail{start: 2, base: '으', pieces: geol}) {
		t.Fatal("으 out of 을 is the link vowel: 을걸 reads as markers from it")
	}
	two := wordsScan("곰가을걸", layerCost)
	two.marks[4] = map[int]bool{1: true, 2: true}
	if two.markerStem(1, 4, tail{start: 3, base: '으', pieces: geol}) {
		t.Fatal("a fused stem of two syllables is judged by its first letter as written (가)")
	}
	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	got := strings.Join(r.Lines("먹었을걸", false), "\n")
	if strings.Contains(got, " 으+(ㄹ)걸") || !strings.Contains(got, "먹[열쇠:동사]+었[선어말:었]+을걸[맺음:을걸]") {
		t.Fatal("먹었을걸\n" + got)
	}
	for _, line := range r.Lines("사과을걸", false) {
		if strings.Contains(line, " 으+(ㄹ)걸") {
			return
		}
	}
	t.Fatal("사과을걸: where 을걸 does not read as markers the new word 으 stands")
}
