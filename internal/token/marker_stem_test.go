//ff:func feature=token type=rule control=sequence
//ff:what TestMarkerStem — a new word after a boundary whose letters read as markers there, or that starts with the link vowel 으 and goes on as markers, is no stem; the unit's first word, a stem whose letters are not so read, a word the table knows, a stem a piece is fused into and a tail that names its stem are let be; so is every stem when there is no word table
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
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
	eu := wordsScan("곰으셨다", layerCost)
	da := tail{start: 3, pieces: []piece{{text: "다", names: "맺음:다", verdict: 1, tail: rule.TailPred}}}
	if eu.markerStem(1, 4, da) {
		t.Fatal("으셨 before 다 when nothing reads 셨다 as markers")
	}
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
	ws := wordsScan("곰서로", layerCost)
	ws.marks[2] = map[int]bool{1: true}
	if ws.markerStem(1, 3, tail{start: 2, pieces: []piece{{text: "로", names: "부사격:로", verdict: 1, tail: rule.TailCase}}}) {
		t.Fatal("a word the table knows (서 the noun)")
	}
}
