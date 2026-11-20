//ff:func feature=parse type=reader control=sequence
//ff:what TestCrossesChange — a new word over a place where the script changes, whose two sides read as an atom or a known word and as a known word or a run, crosses it: A씨 with 가 after it or as a whole, AB씨 past a place where the script does not change, 3분 after a digit, 학교A with Hangul before the run; not so when the side before the place is neither (곰곰A), when the side after it reads as nothing known (K팝 with no 팝), when the tail names its own stem or stands for a number, nor where the script does not change (곰곰했다); the place right before the stem's last syllable a piece is fused into is inside the stem (git했다: git하)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCrossesChange(t *testing.T) {
	ga := []piece{{text: "가", names: "주격:가", verdict: 1, tail: rule.TailCase}}
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}

	sc := wordsScan("A씨가", layerCost)
	sc.ix.lex["씨"] = noun
	if !sc.crossesChange(0, 3, tail{start: 2, pieces: ga}) {
		t.Fatal("A씨 before 가: the atom A and 씨+가")
	}
	if !sc.crossesChange(0, 2, tail{start: 2}) {
		t.Fatal("A씨 as a whole new word")
	}
	if sc.crossesChange(0, 3, tail{start: 2, only: "A씨", pieces: ga}) {
		t.Fatal("a tail that names its own stem is not asked")
	}
	if sc.crossesChange(0, 2, tail{start: 2, as: "수"}) {
		t.Fatal("a tail that stands for a number is not asked")
	}
	past := wordsScan("git했다", layerCost)
	if !past.crossesChange(0, 5, tail{start: 4, base: '하'}) {
		t.Fatal("git했다: the place right before 했, the stem's last syllable a piece is fused into, is inside the stem git하")
	}
	if wordsScan("곰곰했다", layerCost).crossesChange(0, 4, tail{start: 3, base: '하'}) {
		t.Fatal("곰곰했다: no place where the script changes, the place right before 했 among those asked")
	}

	two := wordsScan("AB씨", layerCost)
	two.ix.lex["씨"] = noun
	if !two.crossesChange(0, 3, tail{start: 3}) {
		t.Fatal("AB씨: the script changes only after B, past the place between A and B")
	}
	digit := wordsScan("3분", layerCost)
	digit.ix.lex["분"] = noun
	if !digit.crossesChange(0, 2, tail{start: 2}) {
		t.Fatal("3분: a run of digits is a run")
	}
	if !wordsScan("학교A", layerCost).crossesChange(0, 3, tail{start: 3}) {
		t.Fatal("학교A: the known word 학교 before the run A")
	}
	if wordsScan("곰곰A", layerCost).crossesChange(0, 3, tail{start: 3}) {
		t.Fatal("곰곰A: 곰곰 is neither an atom nor a known word")
	}
	if wordsScan("K팝", layerCost).crossesChange(0, 2, tail{start: 2}) {
		t.Fatal("K팝: 팝 is no known word")
	}
}
