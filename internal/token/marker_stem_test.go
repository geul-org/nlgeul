//ff:func feature=token type=rule control=sequence
//ff:what TestMarkerStem — a new word after a boundary whose letters read as markers there is no stem; the unit's first word, a stem whose letters are not so read, a word the table knows, a stem a piece is fused into and a tail that names its stem are let be; so is every stem when there is no word table
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestMarkerStem(t *testing.T) {
	if none := scanOf("곰에서"); none.markerStem(1, tail{start: 2, pieces: []piece{{text: "서"}}}) || none.marks[2] != nil {
		t.Fatal("without a word table a new word's stem is not asked")
	}
	sc := wordsScan("곰에서", layerCost)
	sc.marks[2] = map[int]bool{1: true}
	seo := tail{start: 2, pieces: []piece{{text: "서", names: "부사격:서", verdict: 1, tail: rule.TailCase}}}
	if !sc.markerStem(1, seo) {
		t.Fatal("에 after a boundary, read as a marker, before 서")
	}
	if sc.markerStem(0, tail{start: 2, pieces: seo.pieces}) {
		t.Fatal("the unit's first word")
	}
	sc.marks[2] = map[int]bool{}
	if sc.markerStem(1, seo) {
		t.Fatal("letters that do not read as markers")
	}
	sc.marks[2] = map[int]bool{1: true}
	if sc.markerStem(1, tail{start: 2, base: '에', pieces: seo.pieces}) || sc.markerStem(1, tail{start: 2, shown: "에", pieces: seo.pieces}) || sc.markerStem(1, tail{start: 2, as: numName, pieces: seo.pieces}) || sc.markerStem(1, tail{start: 2, only: "에", pieces: seo.pieces}) {
		t.Fatal("a tail with its own stem")
	}
	ws := wordsScan("곰서로", layerCost)
	ws.marks[2] = map[int]bool{1: true}
	if ws.markerStem(1, tail{start: 2, pieces: []piece{{text: "로", names: "부사격:로", verdict: 1, tail: rule.TailCase}}}) {
		t.Fatal("a word the table knows (서 the noun)")
	}
}
