//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceAfter — applies when the end before the piece is none of the ends it attaches to
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceAfter(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "after", Mask: EndVowel | EndL | EndOther}}
	hard, _ := PieceAfter(pieceCtx(Piece{PrevEnd: EndHard}), specs)
	vowel, _ := PieceAfter(pieceCtx(Piece{PrevEnd: EndVowel}), specs)
	other, _ := PieceAfter(pieceCtx(Piece{PrevEnd: EndOther}), specs)
	if hard != 1 || vowel != -1 || other != -1 {
		t.Fatal(hard, vowel, other)
	}
}
