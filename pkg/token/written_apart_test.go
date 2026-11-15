//ff:func feature=token type=reader control=sequence
//ff:what TestWrittenApart — a new word whose stem holds, right after a run of Latin letters that is one atom from its head, a title written onto it (A씨 with 가 after it or as a whole, AB씨 past a place where the script does not change) is written apart; not so when the place is right after a digit (3분), when what stands before it from the head is no atom (가A씨), when the letters after it are no such word (K팝), when the tail names its own stem or stands for a number, nor when the place is the stem's last syllable a piece is fused into
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWrittenApart(t *testing.T) {
	ga := []piece{{text: "가", names: "주격:가", verdict: 1, tail: rule.TailCase}}
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}

	sc := wordsScan("A씨가", layerCost)
	sc.ix.lex["씨"] = noun
	if !sc.writtenApart(0, 3, tail{start: 2, pieces: ga}) {
		t.Fatal("A씨 before 가: 씨, a title of article 48, right after the atom A")
	}
	if !sc.writtenApart(0, 2, tail{start: 2}) {
		t.Fatal("A씨 as a whole new word")
	}
	if sc.writtenApart(0, 3, tail{start: 2, only: "A씨", pieces: ga}) {
		t.Fatal("a tail that names its own stem is not asked")
	}
	if sc.writtenApart(0, 2, tail{start: 2, as: "수"}) {
		t.Fatal("a tail that stands for a number is not asked")
	}
	if sc.writtenApart(0, 2, tail{start: 2, base: '씨', pieces: ga}) {
		t.Fatal("the stem's last syllable a piece is fused into is not inside the stem as written")
	}

	two := wordsScan("AB씨", layerCost)
	two.ix.lex["씨"] = noun
	if !two.writtenApart(0, 3, tail{start: 3}) {
		t.Fatal("AB씨: the script changes only after B, past the place between A and B")
	}

	digit := wordsScan("3분", layerCost)
	digit.ix.lex["분"] = noun
	if digit.writtenApart(0, 2, tail{start: 2}) {
		t.Fatal("3분: a place right after a digit is not asked")
	}

	head := wordsScan("가A씨", layerCost)
	head.ix.lex["씨"] = noun
	if head.writtenApart(0, 3, tail{start: 3}) {
		t.Fatal("가A씨: the place after 가 follows no Latin letter, and 가A is no atom")
	}

	pop := wordsScan("K팝", layerCost)
	pop.ix.lex["팝"] = noun
	if pop.writtenApart(0, 2, tail{start: 2}) {
		t.Fatal("K팝: 팝 is no word the spacing articles write apart")
	}
}
