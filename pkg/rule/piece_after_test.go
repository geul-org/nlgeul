//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceAfter — applies when the syllable before the piece does not meet the piece's condition; a non-Hangul one, or none, meets it
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceAfter(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "after", Mask: AfterOpen}}
	hard, _ := PieceAfter(pieceCtx(Piece{Prev: '각'}), specs)
	vowel, _ := PieceAfter(pieceCtx(Piece{Prev: '가'}), specs)
	other, _ := PieceAfter(pieceCtx(Piece{Prev: 'x'}), specs)
	none, _ := PieceAfter(pieceCtx(Piece{}), specs)
	if hard != 1 || vowel != -1 || other != -1 || none != -1 {
		t.Fatal(hard, vowel, other, none)
	}
}
